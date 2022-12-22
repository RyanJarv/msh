import json
import subprocess


def lambda_handler(events, context):
    response = []
    for event in events:
        cmd = event["cmd"]
        body = json.loads(event['body'])

        proc = subprocess.Popen(cmd, stdin=subprocess.PIPE, stdout=subprocess.PIPE)
        body['Content'] = proc.communicate(input=body['Content'].encode())[0]

        response.append(body)

    # TODO implement
    return response
