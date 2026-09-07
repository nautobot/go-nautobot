#!/usr/bin/env python3

import yaml

SPEC_PATH = '/client/api/openapi.yaml'

def _is_stringy_oneof(s):
    return (
        isinstance(s, dict)
        and isinstance(s.get('oneOf'), list)
        and s['oneOf']
        and all(isinstance(b, dict) and b.get('type') == 'string' for b in s['oneOf'])
    )


def _simplify_string_format_oneof(schema):
    """Collapse string alternatives that differ only by their OpenAPI format.

    OpenAPI Generator emits a oneOf wrapper model even when every alternative
    has the same Go representation. Those wrappers cannot be serialized as
    query parameters by the generated Go client.
    """
    if not _is_stringy_oneof(schema):
        return None

    branches_without_format = [
        {key: value for key, value in branch.items() if key != 'format'}
        for branch in schema['oneOf']
    ]
    if any(branch != branches_without_format[0] for branch in branches_without_format[1:]):
        return None

    simplified = dict(branches_without_format[0])
    simplified.update({key: value for key, value in schema.items() if key != 'oneOf'})

    formats = [branch.get('format') for branch in schema['oneOf']]
    if formats[0] is not None and all(value == formats[0] for value in formats[1:]):
        simplified['format'] = formats[0]

    return simplified


def _simplify_string_enum_oneof(schema, schemas):
    """Merge a oneOf of disjoint string enums into one string enum."""
    if (
        not isinstance(schema, dict)
        or not isinstance(schema.get('oneOf'), list)
        or not schema['oneOf']
    ):
        return None

    enum_values = []
    for branch in schema['oneOf']:
        if not isinstance(branch, dict):
            return None

        if '$ref' in branch:
            if set(branch) != {'$ref'}:
                return None
            ref_prefix = '#/components/schemas/'
            ref = branch.get('$ref', '')
            if not ref.startswith(ref_prefix):
                return None
            branch = schemas.get(ref[len(ref_prefix):])

        if (
            not isinstance(branch, dict)
            or set(branch) - {'type', 'enum'}
            or branch.get('type', 'string') != 'string'
            or not isinstance(branch.get('enum'), list)
            or not branch['enum']
            or not all(isinstance(value, str) for value in branch['enum'])
        ):
            return None

        # Overlapping alternatives would match more than one branch, so merging
        # them would change oneOf validation semantics.
        if any(value in enum_values for value in branch['enum']):
            return None
        enum_values.extend(branch['enum'])

    simplified = {'type': 'string', 'enum': enum_values}
    simplified.update({key: value for key, value in schema.items() if key != 'oneOf'})
    return simplified


def _move_object_default_to_properties(schema):
    """Move an object-valued default to defaults on its child properties."""
    if (
        not isinstance(schema, dict)
        or schema.get('type') != 'object'
        or not isinstance(schema.get('properties'), dict)
        or not isinstance(schema.get('default'), dict)
        or not schema['default']
    ):
        return False

    for property_name, default in schema['default'].items():
        property_schema = schema['properties'].get(property_name)
        if not isinstance(property_schema, dict):
            return False
        if 'default' in property_schema and property_schema['default'] != default:
            return False

    for property_name, default in schema['default'].items():
        schema['properties'][property_name]['default'] = default
    schema.pop('default')
    return True

with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():

        # OpenAPI Generator turns "802.1Q Mode" into the invalid Go identifier "8021QMode".
        if name in (
            'PatchedWritableInterfaceRequest',
            'PatchedWritableVMInterfaceRequest',
            'WritableInterfaceRequest',
            'WritableVMInterfaceRequest',
        ):
            mode_property = schema.get('properties', {}).get('mode', {})
            if mode_property.get('title') == '802.1Q Mode':
                print(f"Replacing Go-unsafe title in {name}.properties.mode")
                mode_property['title'] = 'IEEE802.1Q Mode'

        # Remove *_count from required (https://github.com/nautobot/nautobot/issues/6183)
        if 'required' in schema:
            required_fields = schema['required']
            fields_to_remove = [field for field in required_fields if field.endswith('_count')]
            if fields_to_remove:
                print(f"Removing {fields_to_remove} from {name}.required")
                for field in fields_to_remove:
                    required_fields.remove(field)

        # This is intentionally VLAN-specific. JobResult has the same read-only,
        # required computed_fields shape, but only VLAN is known to omit it from
        # API responses (https://github.com/nautobot/nautobot/issues/8082).
        if name == 'VLAN' and 'required' in schema:
            if 'computed_fields' in schema['required']:
                print("Removing computed_fields from VLAN.required")
                schema['required'].remove('computed_fields')

        if 'properties' in schema:
            for property_name, property_schema in schema['properties'].items():
                simplified = _simplify_string_format_oneof(property_schema)
                if simplified is not None:
                    print(
                        f"Simplifying {name}.properties.{property_name} "
                        "string-format oneOf -> string"
                    )
                    schema['properties'][property_name] = simplified
                    property_schema = simplified

                simplified = _simplify_string_enum_oneof(
                    property_schema,
                    data['components']['schemas'],
                )
                if simplified is not None:
                    print(
                        f"Simplifying {name}.properties.{property_name} "
                        "string-enum oneOf -> string enum"
                    )
                    schema['properties'][property_name] = simplified
                    property_schema = simplified

                if _move_object_default_to_properties(property_schema):
                    print(
                        f"Moving {name}.properties.{property_name} object default "
                        "to child properties"
                    )

                # Nullable binary fields generate invalid Go such as
                # "Nullable*os.File".
                # https://github.com/OpenAPITools/openapi-generator/issues/18006
                if (
                    isinstance(property_schema, dict)
                    and property_schema.get('format') == 'binary'
                    and 'nullable' in property_schema
                ):
                    print(f"Removing nullable from binary {name}.properties.{property_name}")
                    property_schema.pop('nullable')

# Traverse path operations
if 'paths' in data:
    # A oneOf of differently formatted strings generates a wrapper model that
    # the Go client's query encoder renders as "<type> value". All alternatives
    # have the same wire representation, so use a plain string parameter.
    for path_name, path_item in data['paths'].items():
        if not isinstance(path_item, dict):
            continue
        for operation_name, operation in path_item.items():
            if not isinstance(operation, dict):
                continue
            for parameter in operation.get('parameters', []):
                parameter_schema = parameter.get('schema')
                simplified = _simplify_string_format_oneof(parameter_schema)
                if simplified is not None:
                    print(
                        f"Simplifying {operation_name.upper()} {path_name} "
                        f"parameter {parameter.get('name')} string-format oneOf -> string"
                    )
                    parameter['schema'] = simplified

    # OpenAPI Generator also uses these inline query schema titles as Go type names.
    for path_name in ('/dcim/interfaces/', '/virtualization/interfaces/'):
        for parameter in data['paths'].get(path_name, {}).get('get', {}).get('parameters', []):
            items = parameter.get('schema', {}).get('items', {})
            if parameter.get('name') == 'mode' and items.get('title') == '802.1Q Mode':
                print(f"Replacing Go-unsafe title in {path_name} GET mode parameter")
                items['title'] = 'IEEE802.1Q Mode'

    # Patch to use AvailableIP array directly instead of PaginatedAvailableIPList
    # (https://github.com/nautobot/nautobot/issues/2131)
    if '/ipam/prefixes/{id}/available-ips/' in data['paths']:
        available_ips_path = data['paths']['/ipam/prefixes/{id}/available-ips/']
        if 'get' in available_ips_path and 'responses' in available_ips_path['get']:
            responses = available_ips_path['get']['responses']
            if '200' in responses and 'content' in responses['200']:
                print("Updating available-ips GET response to return an array of AvailableIP objects")
                responses['200']['content']['application/json']['schema'] = {
                    'type': 'array',
                    'items': {'$ref': '#/components/schemas/AvailableIP'}
                }
                responses['200']['content']['text/csv']['schema'] = {
                    'type': 'array',
                    'items': {'$ref': '#/components/schemas/AvailableIP'}
                }
        if 'post' in available_ips_path and 'responses' in available_ips_path['post']:
            responses_post = available_ips_path['post']['responses']
            if '201' in responses_post and 'content' in responses_post['201']:
                print("Updating available-ips POST response to return an array of IPAddress objects")
                responses_post['201']['content']['application/json']['schema'] = {
                    'type': 'array',
                    'items': {'$ref': '#/components/schemas/IPAddress'}
                }
                responses_post['201']['content']['text/csv']['schema'] = {
                    'type': 'array',
                    'items': {'$ref': '#/components/schemas/IPAddress'}
                }

with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
