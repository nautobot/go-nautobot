#!/usr/bin/env python3

import yaml

SPEC_PATH = '/client/api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():
    
        # Handle Manufacturer schema (https://github.com/nautobot/nautobot/issues/6183)
        if name == 'Manufacturer' and 'required' in schema:
            required_fields = schema['required']
            if 'cloud_account_count' in required_fields:
                print(f"Removing 'cloud_account_count' from {name}.required")
                required_fields.remove('cloud_account_count')
                
        # Handle failover_strategy
        if 'failover_strategy' in schema.get('properties', {}):
            prop = schema['properties']['failover_strategy']
            if 'default' in prop and prop['default'] is None:
                print(f"Removing 'default: null' in {name}.failover_strategy")
                del prop['default']
                
        # Handle PowerFeed schema
        if name == 'PowerFeed' and 'properties' in schema:
            # Replace nested `type` field in PowerFeed
            if 'type' in schema['properties']:
                type_property = schema['properties']['type']
                if 'properties' in type_property and 'value' in type_property['properties']:
                    value_property = type_property['properties']['value']
                    if 'enum' in value_property and set(value_property['enum']) == {'primary', 'redundant'}:
                        print(f"Replacing complex 'type' field in PowerFeed")
                        schema['properties']['type'] = {
                            'type': 'string',
                            'enum': ['primary', 'redundant'],
                            'default': 'primary'
                        }

            # Replace nested `supply` field in PowerFeed
            if 'supply' in schema['properties']:
                supply_property = schema['properties']['supply']
                if 'properties' in supply_property and 'value' in supply_property['properties']:
                    value_property = supply_property['properties']['value']
                    if 'enum' in value_property and set(value_property['enum']) == {'ac', 'dc'}:
                        print(f"Replacing complex 'supply' field in PowerFeed")
                        schema['properties']['supply'] = {
                            'type': 'string',
                            'enum': ['ac', 'dc'],
                            'default': 'ac'
                        }

            # Replace nested `phase` field in PowerFeed
            if 'phase' in schema['properties']:
                phase_property = schema['properties']['phase']
                if 'properties' in phase_property and 'value' in phase_property['properties']:
                    value_property = phase_property['properties']['value']
                    if 'enum' in value_property and set(value_property['enum']) == {'single-phase', 'three-phase'}:
                        print(f"Replacing complex 'phase' field in PowerFeed")
                        schema['properties']['phase'] = {
                            'type': 'string',
                            'enum': ['single-phase', 'three-phase'],
                            'default': 'single-phase'
                        }
        
        # Handle Prefix schema
        if name == 'Prefix' and 'properties' in schema:
            # Replace complex `type` object with a simpler string enum in Prefix
            if 'type' in schema['properties']:
                type_property = schema['properties']['type']
                if 'properties' in type_property and 'value' in type_property['properties']:
                    print(f"Replacing complex 'type' field in Prefix")
                    schema['properties']['type'] = {
                        'type': 'string',
                        'enum': ['container', 'network', 'pool'],
                        'default': 'network'
                    }
                    
        if 'properties' in schema:
            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]

            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    if schema['properties'][ntype]['format'] == 'binary':
                        schema['properties'][ntype].pop('nullable')

# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
