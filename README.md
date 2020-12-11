## [WIP] Multivac Shell

This is really a mix of several ideas I have that may or may not be related. Please don't actually use this project, 
it's really just for experimentation at the moment.

See the [v1 branch](https://github.com/RyanJarv/msh/tree/v1) for a list of (likely incoherent) notes related to this.

### Setup
```
git clone https://github.com/RyanJarv/msh.git
cd msh
make install
```

### Dockerfile Examples

We will always assume the context directory is where the Dockerfile is stored. Use the docker-compose support if you 
want to change this.

#### Local Execution
```
% cat ./test/echo.msh
#!/usr/local/bin/msh dockerfile
FROM alpine:3
ENTRYPOINT ["echo"]
```

```
% ./test/echo.msh hello world
2020/11/23 22:31:12 Running command and waiting for it to finish...
[+] Building 0.0s (5/5) FINISHED                                                                                                                                                               
 => [internal] load build definition from echo.msh                                                                                                                                            
 => => transferring dockerfile: 191B                                                                                                                                                          
 => [internal] load .dockerignore                                                                                                                                                             
 => => transferring context: 2B                                                                                                                                                               
 => [internal] load metadata for docker.io/library/alpine:3                                                                                                                                   
 => CACHED [1/1] FROM docker.io/library/alpine:3                                                                                                                                              
 => exporting to image                                                                                                                                                                        
 => => exporting layers                                                                                                                                                                       
 => => writing image sha256:595b26ac4285e518a9794f0f8287d6a3fb2df9dfeb636d11a9ca5fd8070424b4                                                                                                  
 => => naming to docker.io/library/echo.msh                                                                                                                                                   
2020/11/23 22:31:12 Running command and waiting for it to finish...
hello world
```

#### ECS Remote Execution

This assumes you have a default ECS cluster set up. If you don't you can create run by running `aws ecs 
create-cluster --cluster-name default`.

```
% msh ecs ./test/echo.msh hello from ecs
2020/11/24 02:22:06 Running command and waiting for it to finish...
2020/11/24 02:22:06 Running command and waiting for it to finish...
✔ Successfully provisioned task resources.

[+] Building 0.0s (5/5) FINISHED                                                                                                                                                               
 => [internal] load build definition from echo.msh                                                                                                                                        0.0s
 => => transferring dockerfile: 34B                                                                                                                                                       0.0s
 => [internal] load .dockerignore                                                                                                                                                         0.0s
 => => transferring context: 2B                                                                                                                                                           0.0s
 => [internal] load metadata for docker.io/library/alpine:3                                                                                                                               0.0s
 => CACHED [1/1] FROM docker.io/library/alpine:3                                                                                                                                          0.0s
 => exporting to image                                                                                                                                                                    0.0s
 => => exporting layers                                                                                                                                                                   0.0s
 => => writing image sha256:595b26ac4285e518a9794f0f8287d6a3fb2df9dfeb636d11a9ca5fd8070424b4                                                                                              0.0s
 => => naming to 253528964770.dkr.ecr.us-east-2.amazonaws.com/copilot-msh:latest                                                                                                          0.0s
Login Succeeded
The push refers to repository [253528964770.dkr.ecr.us-east-2.amazonaws.com/copilot-msh]
ace0eda3e3be: Layer already exists 
latest: digest: sha256:0b5ed5d28d9a085a56a49f6fc1e0740e8e16f02a8cffaf4279d0247e47cca10b size: 527
✔ Successfully updated image to task.

✔ Task msh is running.

copilot-task/msh/432917be hello from ecs
Task has stopped.
```

#### Using it Like a Normal Dockerfile

What's cool here is the shebang line is just a comment to docker, so you can treat it like a normal Dockerfile.

```
% docker build -t test -f test/echo.msh test
[+] Building 0.1s (5/5) FINISHED                                                                                                                                                                              
 => [internal] load build definition from Dockerfile                                                                                                                                                     0.0s
 => => transferring dockerfile: 108B                                                                                                                                                                     0.0s
 => [internal] load .dockerignore                                                                                                                                                                        0.0s
 => => transferring context: 2B                                                                                                                                                                          0.0s
 => [internal] load metadata for docker.io/library/alpine:3                                                                                                                                              0.0s
 => CACHED [1/1] FROM docker.io/library/alpine:3                                                                                                                                                         0.0s
 => exporting to image                                                                                                                                                                                   0.0s
 => => exporting layers                                                                                                                                                                                  0.0s
 => => writing image sha256:595b26ac4285e518a9794f0f8287d6a3fb2df9dfeb636d11a9ca5fd8070424b4                                                                                                             0.0s
 => => naming to docker.io/library/test
```

### docker-compose

#### Local

docker-compose works similarly but when executed behaves as if "docker-compose run --service-ports app" was run against it.

```yaml
#!/usr/bin/env msh compose
version: "3.8"
services:
  app:
    image: alpine:3
    entrypoint: echo
```

#### Local Lambda Compose

I'm re-thinking how lambda support for this works at the moment but you can find an example of running lambda locally in [./test/compose/lambda](./test/compose/lambda).

Start the lambda container
```
./test/compose/lambda/compose.yaml app.lambda_handler 
```

In another shell trigger an invocation
```
curl  -XPOST 'http://localhost:9000/2015-03-31/functions/function/invocations' -d '{}'
```

Note: The test app is set up to exit after one run so that it can be tested with `make compose/lambda`


### Scripts

You can simply run the ./test/ruby.rb and it will behave the same way as the local lambda compose example. I'll likely be changing this behavior to convert cmd args to a single request and then exiting. This behavior will likely be used for local testing where args map in some way to API Gateway path/parameters.
