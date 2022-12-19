# Custom AWS Lambda runtime: provide your own bootstrap

A runtime is a program that runs a Lambda function's handler method when the function is invoked. The runtime can
be included in your function's deployment package, or in a [layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
[Learn more](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-custom.html) about custom Lambda runtimes.

For reference, this function includes a sample [Bash](https://www.gnu.org/software/bash/) runtime in `bootstrap.sample` and a corresponding
handler file `hello.sh.sample`. As a next step, you should provide your own bootstrap by either adding a layer implementing a custom runtime or
including a `bootstrap` file in this function's deployment package.
