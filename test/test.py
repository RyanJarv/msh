#!/usr/bin/env msh.lambda
import json

 af def   lambda_handler  (event, context):
    print(event)
    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }

