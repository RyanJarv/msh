import json


def lambda_handler(events, context):
    print("event: " + json.dumps(events))
    print("context: " + json.dumps(events))

    response = []
    for event in events:
        body = json.loads(event['body'])
        body['Content'] = 'modified\n'
        response.append(body)
        # response.append(json.dumps(body))

    # TODO implement
    return response
