## Multivac Shell

Note: This is really a mix of several ideas I have that may or may not be related.

Blog post: https://blog.ryanjarv.sh/2020/11/22/msh.html


### Other

Previous (somewhat) related project: https://github.com/RyanJarv/coderun


### Examples

In case you read my post or notes in the coderun repo, it might be confusing how this will all fit together.. I'm still working on that part but this is at least what is working so far.

## Echo
```
% cat echo.msh                              
#!/usr/local/bin/msh exec docker build -t {{.Name}} -f {{.Path}} {{.ContextDir}}; docker run {{.Name}} hello world
FROM alpine:3
ENTRYPOINT ["echo"]
```

```
% ./echo.msh                               
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

## Cat Example
The `dockerfile` arg is mostly doing the same thing as the shebang line in the echo example. The idea is the syntax would be closer to this example but could be written out manually as well if you needed to.

```
% cat cat.msh                               
#!/usr/local/bin/msh dockerfile
FROM alpine:3
ENTRYPOINT ["cat"]
```

```
% echo 'hello world again'| ./cat.msh                               
2020/11/23 22:31:43 Running command and waiting for it to finish...
[+] Building 0.0s (5/5) FINISHED                                                                                                                                                               
 => [internal] load build definition from cat.msh                                                                                                                                             
 => => transferring dockerfile: 104B                                                                                                                                                          
 => [internal] load .dockerignore                                                                                                                                                             
 => => transferring context: 2B                                                                                                                                                               
 => [internal] load metadata for docker.io/library/alpine:3                                                                                                                                   
 => CACHED [1/1] FROM docker.io/library/alpine:3                                                                                                                                              
 => exporting to image                                                                                                                                                                        
 => => exporting layers                                                                                                                                                                       
 => => writing image sha256:a71009c164801bcd6572e22ca1c86da33a90dcb78c1b75c1187f74d62e0eb8a6                                                                                                  
 => => naming to docker.io/library/cat.msh                                                                                                                                                    
2020/11/23 22:31:43 Running command and waiting for it to finish...
hello world again
```

