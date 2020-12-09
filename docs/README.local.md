# WIP Concepts

## Invocation Example
```bash
% msh new ProjectName
[ProjectName] % ./app.rb
[ProjectName] % msh remote aws dev
[ProjectName aws dev] % ./app.rb
[ProjectName aws dev] % exit
[ProjectName] % exit
% exit
```

## Config Pipeline

Contains configuration about the executing pipeline, values like ProjectName and StdoutUri are handled by msh to make
streaming between environments seamless.

echo hello | msh aws -- ./app.rb
    msh handles:
        * setting up aws environment configuration
        * executing ./app.rb
    app.rb also has '#!/usr/bin/env msh lambda ruby' which:
        * deploying shell script as a lambda ruby function
        * stdin lines to lambda calls with APIGateway context
        * streams output from lambda to stdout
msh aws -- ./app1.rb | msh aws -- ./app.rb
    * msh for app1 detects stdout is not the terminal
        * local pipe contains structured info with stdout info



        * lines to lambda calls
    
        * client invocation with APIGateway context (doesn't actually use api gateway, just context structure passed)
        * StdoutURI set to APIGateway 
        * Current ARN passed to next command for IAM permissions if needed
lambda stdout to lambda: client invocation with http context
    same
    

Taking lines of input from external tools we convert it to the internal format. A nil value indicates EOF.
```go
stdout := Line{
    Config: &Config{
        ProjectName: "MyApp",
        StdoutUri:   "sqs queue",
    },
}
```


```
ProjectName % { while :; do sleep 1; echo "hello"; } |              ./app       |                msh remote ./app
                                                     | init then lines to json  | deploy if needed then json to http then destroy on EOF
```

```bash
                    msh remote ./app                             |                 
deploy if needed then stdin lines to http then destroy stdin EOF |
```

## Local Invocation Example
```bash
% cat ../msh-repo/conf/ruby/app.rb|head -n 1
#!/usr/local/msh/msh lambda ruby
%
% ./app.rb
  <runs in lib>
```
        

