#!/usr/bin/env python3
import itertools
import json
import sys
from copy import copy
from pathlib import Path

data = json.loads(Path('schema/dynamic-config.json').read_text())

events = {}
section = data['default']
for svc, svc_data in section['services'].items():
    for event_name, event_data in svc_data['eventTypes'].items():
        event = {
            "source": [svc],
        }

        print(json.dumps(event))

        if 'patternPaths' not in event_data:
            print(json.dumps(event))
            continue

        for v in event_data['patternPaths']:
            if len(v['selector']) != 1:
                print(f'unknown selector: {svc} {event_name} {v} {v["selector"]}', file=sys.stderr)
                continue

            event[v['selector'][0]] = v['values']
            print(json.dumps(event))

        # if event_data['eventTypeTemplate'] != 'AWSAPIType':
        #     print(f'unknown template: {svc} {event_data["eventTypeTemplate"]}', file=sys.stderr)
        #     continue

        value_options = []
        for component_name, component in event_data['components'].items():
            if component['type'] != 'MULTI_SELECT_COMPONENT':
                continue
            elif type(component['options']) is list:
                # TODO: look into this
                continue
            elif component['options'].get('type', None) != 'INTERNAL_RESOURCES':
                continue

            values = section
            for path in component['options']['selector']:
                values = values[path]

            value_options.append([{
                "value": v,
                "location": component['selector'],
            } for v in values])

        if not value_options:
            continue

        permutations = itertools.product(*value_options)

        new_event = copy(event)

        for opt_set in permutations:
            for opt in opt_set:
                curr = new_event
                for name, i in enumerate(opt['location']):
                    if len(opt['location']) == i:
                        for v in values:
                            curr = v
                    else:
                        curr = {}

            print(json.dumps(new_event))