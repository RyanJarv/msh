## The Multivac Shell

***Important: This project is not ready for production.***

The Multivac Shell allows you to deploy simple pipelines using CloudFormation (via CDK) expressed through functional programming 
(shell pipes, map, filter, etc...) in the same shell you know and love.

![F94zTUBbwAA51JW](https://github.com/RyanJarv/msh/assets/4079939/8c120edf-35b9-48da-ad1c-fcef7f62fb87)

Above we can see the Multivac Shell used to deploy a cron job that runs every 5 minutes, which then triggers a state 
machine which:

* Describes all instances in the account.
* Filters out the instances that have been running less than 72 hours.
* Shuts down the remaining instances, the one's that when't filtered previously. 

Once deployed this pipeline is persistent in the account and must be manually deleted in the CloudFormation console to
deactivate it.

# What Is Happening Here?

So... good question, this is actually running in a real shell, and we aren't doing anything too tricky with [overriding shell 
syntax](./bin/env.sh).

We're using CDK in generally how it's intended to, however we're expressing the configuration through piped commands.
To do this we build a simple internal configuration consisting of the arguments passed to each function. This 
configuration is read in through the first line of stdin on startup, and printed as the first line to stdout, but only 
if stdin/stdout is not a TTY.

You can see this internal configuration by appending cat or jq to the end of the command above:

```
(msh) me@Ryans-MacBook-Pro msh % @cron '0 0 * * * *' | sfn{ 
        .aws ec2 describe-instances --query 'Reservations[]. Instances[]' \
        | .foreach \
        | .filter ./bin/hours_running.py 72
}  | jq
{
  "Steps": [
    {
      "Name": "schedule",
      "Value": { "Rule": null, "CronOptions": { "day": null, "hour": "0", "minute": "0", "month": "*", "weekDay": null, "year": "*" } }
    },
    {
      "Name": "sfn",
      "Value": { "Name": "sfn" }
    },
    {
      "Name": "aws",
      "Value": {
        "LambdaInvoke": null,
        "Script": "#!/usr/bin/env python3\n# Borrowed from https://gist.github.com/dagrz/9b006d116302ba9dc4a1234b867c30c7, thanks dagrz!\n\nimport json\nimport sys\nimport awscli.clidriver\nfrom io import StringIO\n\ndriver = awscli.clidriver.create_clidriver()\n\n\ndef lambda_handler(event, context):\n    old_stdout = sys.stdout\n    sys.stdout = mystdout = StringIO()\n\n    driver.main(args=event['command'])\n\n    sys.stdout = old_stdout\n\n    mystdout.seek(0)\n\n    return json.loads(mystdout.read())\n\n\nif __name__ == '__main__':\n    lambda_handler({'command': sys.argv[1:]}, {})\n    sys.exit(0)\n",
        "Args": [ ".aws", "ec2", "describe-instances", "--query", "Reservations[]. Instances[]" ],
        "IamStatementProps": [ { "actions": [ "ec2:DescribeInstances" ], "conditions": null, "effect": "", "notActions": null, "notPrincipals": null, "notResources": null, "principals": null, "resources": [ "*" ], "sid": null } ],
        "Environment": { "PYTHONPATH": "/opt/awscli" }
      }
    },
    {
      "Name": "map",
      "Value": { "IChain": null }
    },
    {
      "Name": "filter",
      "Value": {
        "INextable": null,
        "Start": null,
        "Lambda": { "IChain": null, "Function": null, "LambdaOpts": { "InputPath": "$.__input", "ResultSelector": { "result.$": "$.Payload" }, "OutputPath": null, "ResultPath": "$.__choice" },
          "Script": "#!/usr/bin/env .filter\n\nfrom datetime import datetime, timezone\nimport os\n\n\ndef lambda_handler(event, context):\n    if event['State']['Name'] != 'running':\n        return False\n\n    input_date_str = event['UsageOperationUpdateTime']\n    print(input_date_str)\n    input_date = datetime.fromisoformat(input_date_str)\n\n    current_date = datetime.now(timezone.utc)\n    delta = current_date - input_date\n\n    return (delta.total_seconds() / 60 / 60) > int(os.environ['ARG1'])",
          "Args": [ ".filter", "./bin/hours_running.py", "72" ],
          "Environment": { "ARG0": "./bin/hours_running.py", "ARG1": "72" },
          "TimeoutSeconds": 300
        }
      }
    }
  ]
}

```



JSON, and


Multivac Shell Commands do not behave like most other commands. 

your used to. 




## Challange 1: Write a fork bomb for AWS in the Multivac Shell

Request: https://x.com/eigenseries/status/1719883102187876768


Result: https://www.youtube.com/watch?v=fBQqqTbr6zw

## More Fun Ideas?

Feel free to open an issue.

## Boto3 IAM Mappings

For those who are interested in the AWS CLI IAM mappings of IAN's [iam-dataset](https://github.com/iann0036/iam-dataset)
you can find them in this [map file](./data/map.json). The code that generates them is in [scripts](./scripts/awscli_iam_map).

