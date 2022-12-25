import json
import subprocess
import os
import base64


def lambda_handler(events, context):
    print("events: " + json.dumps(events))
    print("context: " + json.dumps(context, default=str))

    response = []
    for event in events:
        # Body will be a json string, so decode again.
        msg = json.loads(event['body'])

        env = event["env"]
        cmd = event["cmd"]

        print("body (decoded): " + str(msg))
        print("cmd env vars: " + json.dumps(env))
        print("running cmd: " + ' '.join(cmd))

        if msg['Type'] != 1:
            # Just forward the event if type is not message.
            response.append(event['body'])
            continue

        os.environ.update(env)

        proc = subprocess.Popen(cmd, stdin=subprocess.PIPE, stdout=subprocess.PIPE)

        output = proc.communicate(input=msg['Content'].encode())[0]

        # Encode twice so that output is a json encoded string.
        msg['Content'] = output.decode()

        print("output: " + str(msg))

        msg = json.dumps(msg)
        response.append(msg)

    # TODO implement
    return response
