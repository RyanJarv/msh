## Multivac Shell

Note: This is really a mix of several ideas I have that may or may not be related.

Blog post: https://blog.ryanjarv.sh/2020/11/22/msh.html


### Other

Previous (somewhat) related project: https://github.com/RyanJarv/coderun


### Examples

In case you read my post or notes in the coderun repo, it might be confusing how this will all fit together.. I'm still working on that part but this is at least what is working so far.

#### Echo
```
% cat echo.msh                              
#!/usr/local/bin/msh dockerfile
FROM alpine:3
ENTRYPOINT ["echo"]
```

```
% ./echo.msh hello world
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

#### ECS

This assumes you have a default ECS cluster set up. If you don't you can create run by running `aws ecs create-cluster --cluster-name default`.

```
% msh ecs ./config/echo.msh hello from ecs
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

