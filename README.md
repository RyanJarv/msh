## The Multivac Shell

***Important: This project is not ready for production and only supports a single stack currently.***

The Multivac Shell allows you to deploy simple pipelines using CloudFormation (via CDK) expressed through functional programming 
(shell pipes, map, filter, etc...) in the same shell you know and love.

![Screenshot 2024-08-15 at 6 56 37 PM](https://github.com/user-attachments/assets/0d730d27-cc99-4bf3-8ad4-d0c68787e8f5)

Above we can see the Multivac Shell used to deploy a cron job that runs every 5 minutes, which then triggers a state 
machine which:

* Describes all instances in the account.
* Filters out the instances that have been running less than 72 hours.
* Shuts down the remaining instances, the one's that when't filtered previously. 

Once deployed this pipeline is persistent in the account and must be manually deleted in the CloudFormation console to
deactivate it.

## Getting Started

Ensure golang is installed, then run the following commands:

```bash
git clone git@github.com:RyanJarv/msh.git
cd msh
make pull
make build
. ./bin/env.sh
```

You should now be able to deploy the following pipeline:
```bash
@cron '0 0 * * * *' | sfn{ 
        .aws ec2 describe-instances --query 'Reservations[]. Instances[]' \
        | .foreach \
        | .filter ./bin/hours_running 72 \
        | .aws ec2 stop-instances --instance-ids '$.Instanceld'
}
```

## Interesting Features

* Easy deploy lambda's by running the code file directly.
  * Running [./scripts/hello_world.py](./scripts/hello_world.py), will deploy a lambda function that returns "Hello World!".
* Automatic IAM policy generation using a fork of Ian's iam-dataset [built for the awscli](cmd/aws/cli_iam_map.json).
* [Experimental support](https://github.com/RyanJarv/msh/tree/main/cmd.experimental/api) for compiling botocore paginator 
definitions to CDK step function definitions.
* Easily add custom lambda commands by taking advantage of [shebang lines](https://github.com/RyanJarv/msh/blob/d0675cb3195d18bbca4e20cbecd080e4c0c5e033/bin/hours_running.py#L1).
  * `.filter ./bin/hours_running 72` can also be run as `hours_running 72`.

## What Is Happening Here?

So... good question, this is actually running in a real shell, and we aren't doing anything too tricky with [overriding shell 
syntax](./bin/env.sh).

We're using CDK in generally how it's intended to, however we're expressing the configuration through piped commands.
To do this we build a [simple internal configuration](https://github.com/RyanJarv/msh/blob/9ad244708619a73d14ef61bdfbd86edf0f70db4e/pkg/app/app.go#L49)
consisting of the arguments passed to each function. This configuration is [read in through the first line of stdin](https://github.com/RyanJarv/msh/blob/e897079a8af68f4ec99d9ca99545b9c470a0ea5f/pkg/app/app.go#L29)on startup, extended,
then [printed as the first line to stdout](https://github.com/RyanJarv/msh/blob/e897079a8af68f4ec99d9ca99545b9c470a0ea5f/pkg/app/run.go#L25), 
but only if stdin/stdout is not a TTY.

You can see this internal configuration by appending cat or jq to the end of the command above:

```
(msh) me@Ryans-MacBook-Pro msh % @cron '0 0 * * * *' | sfn{ 
        .aws ec2 describe-instances --query 'Reservations[].Instances[]' \
        | .foreach \
        | .filter ./bin/hours_running 72
}  | jq
{
  "Steps": [
    { "Name": "schedule", "Value": { "CronOptions": { "day": null, "hour": "0", "minute": "0", "month": "*", "weekDay": null, "year": "*" } } },
    { "Name": "sfn",      "Value": {} },
    { "Name": "aws",      "Value": { "Script": "...", "Args": [".aws", "ec2", "describe-instances", ...], "IamStatementProps": [...], "Environment": {...} } },
    { "Name": "map",      "Value": {}},
    { "Name": "filter",   "Value": { "Args": [...], "Opts": {...} } },
    { "Name": "aws",      "Value": { "Script": "...", "Args": [".aws", "ec2", "stop-instances", ...], "IamStatementProps": [...], "Environment": {...} } }
  ]
}
```

Each step value represents a marshalled [CdkStep](https://github.com/RyanJarv/msh/blob/d0675cb3195d18bbca4e20cbecd080e4c0c5e033/cmd/filter/filter.go#L29)
after the [New](https://github.com/RyanJarv/msh/blob/d0675cb3195d18bbca4e20cbecd080e4c0c5e033/cmd/filter/filter.go#L13C6-L13C9) function is called.

If the current executing [command is a TTY](https://github.com/RyanJarv/msh/blob/d0675cb3195d18bbca4e20cbecd080e4c0c5e033/pkg/app/run.go#L29) we 
finish building the CDK state by calling Compile on each CdkStep in the configuration, [iterate through the reversed steps](https://github.com/RyanJarv/msh/blob/d0675cb3195d18bbca4e20cbecd080e4c0c5e033/pkg/app/run.go#L66)
to build the chain from the end to the beginning. The beginning of the chain is extended in different ways 
depending on the underlying type, whether it is an awsstepfunctions.INextable, awsevents.Rule, awssns.ITopic, or 
types.IIterator.

# Develop Your Own Steps

Developing a step is fairly simple, just add a new directory under [./cmd](./cmd) which contains the following code:

```golang
// New returns the CdkStep with any configuration needed to build the required CDK resources.
func New(app app.App) (*Cmd, error) {
	return &Cmd{}, nil
}

// Cmd is a struct that implements types.CdkStep
//
// This represents the internal configuration of the step, and is passed between the 
// steps using JSON. You don't need to worry about this however, the struct will be 
// restored for you before Compile is called.
//
// Cmd is passed to the preceding step in the pipeline, so you may need to ensure the 
// correct interfaces are implemented depending on your use case.
//
// Interfaces representing how Cmd can extend the chain, if none are implemented no 
// step can precede this command:
//   **awsevents.IRuleTarget**
//        Implements an awsevents RuleTarget, Cmd may be passed to the previous step's 
//        AddTarget method.
//   **types.IChainable**
//        Implements a sfn step which can have preceding steps, Cmd may be passed to the 
//        previous step's Next method.
//   **awssns.ITopicSubscription**
//        Implements an SNS topic subscription, Cmd may be passed to the previous step's 
//        AddSubscription method.
//
// Interfaces representing how the chain can be extended, if none are impelmented no step 
// can follow this command:
//   **awsstepfunctions.INextable**
//        Implements a sfn step which can have steps after it, the next step is passed to 
//        our Next method.
//   **types.IIterator**
//        Implements a sfn step which iterates over set of steps, the next step will be
//        passed to our Iterator method.
//   **awsevents.Rule**
//        Implements an awsevents Rule, the next step will be passed to our AddTarget 
//        method.
type Cmd struct {}

// GetName needs to be a unique name for the step, this is used to identify the step in the internal configuration.
func (s Cmd) GetName() string { return "cmd-name" }

// Compile is called when we are building the CDK stack. All CDK resources need to be created here.
func (s *Cmd) Compile(stack constructs.Construct, i int) error {
	return nil
}
```



## Challange 1: Write a fork bomb for AWS in the Multivac Shell

***Note:*** This is not using the current fork, I had to hack with things a bit to get this to work.

Request: https://x.com/eigenseries/status/1719883102187876768


Result: https://www.youtube.com/watch?v=fBQqqTbr6zw

## More Fun Ideas?

Feel free to open an issue.

## Boto3 IAM Mappings

For those who are interested in the AWS CLI IAM mappings of IAN's [iam-dataset](https://github.com/iann0036/iam-dataset)
you can find them in this [map file](./data/map.json). The code that generates them is in [scripts](./scripts/awscli_iam_map).

