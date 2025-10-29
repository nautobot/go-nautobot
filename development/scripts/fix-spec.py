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

with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():

        # Remove *_count from required (https://github.com/nautobot/nautobot/issues/6183)
        if 'required' in schema:
            required_fields = schema['required']
            fields_to_remove = [field for field in required_fields if field.endswith('_count')]
            if fields_to_remove:
                print(f"Removing {fields_to_remove} from {name}.required")
                for field in fields_to_remove:
                    required_fields.remove(field)

        # PowerFeed patches (go bindings generator issue?)
        if name == 'PowerFeed' and 'properties' in schema:
            if 'type' in schema['properties']:
                type_property = schema['properties']['type']
                if 'properties' in type_property and 'value' in type_property['properties']:
                    value_property = type_property['properties']['value']
                    if 'enum' in value_property and set(value_property['enum']) == {'primary', 'redundant'}:
                        print(f"Replacing complex 'type' field in PowerFeed")
                        schema['properties']['type'] = {
                            'type': 'object',
                            'properties': {
                                'value': {'type': 'string','enum': ['primary','redundant'],'default': 'primary'},
                                'label': {'type': 'string','enum': ['Primary','Redundant'],'default': 'Primary'},
                            }
                        }
            if 'supply' in schema['properties']:
                supply_property = schema['properties']['supply']
                if 'properties' in supply_property and 'value' in supply_property['properties']:
                    value_property = supply_property['properties']['value']
                    if 'enum' in value_property and set(value_property['enum']) == {'ac', 'dc'}:
                        print(f"Replacing complex 'supply' field in PowerFeed")
                        schema['properties']['supply'] = {
                            'type': 'object',
                            'properties': {
                                'value': {'type': 'string','enum': ['ac','dc'],'default': 'ac'},
                                'label': {'type': 'string','enum': ['AC','DC'],'default': 'AC'},
                            }
                        }
            if 'phase' in schema['properties']:
                phase_property = schema['properties']['phase']
                if 'properties' in phase_property and 'value' in phase_property['properties']:
                    value_property = phase_property['properties']['value']
                    if 'enum' in value_property and set(value_property['enum']) == {'single-phase', 'three-phase'}:
                        print(f"Replacing complex 'phase' field in PowerFeed")
                        schema['properties']['phase'] = {
                            'type': 'object',
                            'properties': {
                                'value': {'type': 'string','enum': ['single-phase','three-phase'],'default': 'single-phase'},
                                'label': {'type': 'string','enum': ['Single phase','Three-phase'],'default': 'Single phase'},
                            }
                        }

        # Prefix patch (go bindings generator issue?)
        if name == 'Prefix' and 'properties' in schema:
            if 'type' in schema['properties']:
                type_property = schema['properties']['type']
                if 'properties' in type_property and 'value' in type_property['properties']:
                    print(f"Replacing complex 'type' field in Prefix with detailed properties")
                    schema['properties']['type'] = {
                        'type': 'object',
                        'properties': {
                            'value': {'type': 'string','enum': ['container','network','pool'],'default': 'network'},
                            'label': {'type': 'string','enum': ['Container','Network','Pool'],'default': 'Network'},
                        }
                    }

        if 'properties' in schema:
            # Email: oneOf -> plain string with default ""
            email_prop = schema['properties'].get('email')
            if _is_stringy_oneof(email_prop):
                print(f"Simplifying {name}.properties.email oneOf -> string/default")
                simplified = {'type': 'string', 'maxLength': 254, 'default': ''}
                if isinstance(email_prop, dict) and 'title' in email_prop:
                    simplified['title'] = email_prop['title']
                schema['properties']['email'] = simplified

            # failover_strategy: oneOf of enums -> string enum with default (TODO: create go bindings generator issue)
            fs_prop = schema['properties'].get('failover_strategy')
            if isinstance(fs_prop, dict) and 'oneOf' in fs_prop:
                enums = []
                for branch in fs_prop.get('oneOf', []):
                    if '$ref' in branch:
                        refname = branch['$ref'].rsplit('/', 1)[-1]
                        refschema = data['components']['schemas'].get(refname, {})
                        if isinstance(refschema, dict) and 'enum' in refschema:
                            enums.extend(refschema['enum'])
                    elif 'enum' in branch:
                        enums.extend(branch['enum'])
                seen = set()
                enums = [e for e in enums if (e not in seen and not seen.add(e))]
                if enums:
                    print(f"Simplifying {name}.properties.failover_strategy oneOf -> string/enum")
                    schema['properties']['failover_strategy'] = {
                        'type': 'string',
                        'enum': enums,
                        'default': '' if '' in enums else enums[0]
                    }

            # Remove object default for structured failover_strategy (value/label) (TODO: create go bindings generator issue)
            fs_obj = schema['properties'].get('failover_strategy')
            if (
                isinstance(fs_obj, dict)
                and fs_obj.get('type') == 'object'
                and isinstance(fs_obj.get('properties'), dict)
                and 'value' in fs_obj['properties'] and 'label' in fs_obj['properties']
                and isinstance(fs_obj.get('default'), dict)
            ):
                print(f"Removing object default from {name}.properties.failover_strategy")
                fs_obj.pop('default', None)

            # Non-nullable binaries (https://github.com/OpenAPITools/openapi-generator/issues/18006)
            for ntype in ('front_image', 'rear_image'):
                if ntype in schema['properties']:
                    if schema['properties'][ntype].get('format') == 'binary':
                        schema['properties'][ntype].pop('nullable', None)

# Patch to use AvailableIP array directly instead of PaginatedAvailableIPList (https://github.com/nautobot/nautobot/issues/2131)
if 'paths' in data:
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
