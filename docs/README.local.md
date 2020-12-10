# Random Ideas

Really I think this is all a distraction for now, and unsure if it even makes sense. Just consider it a collection of random thoughts.

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

## Config Pipeline for remote pipelining

Contains configuration about the executing pipeline, values like ProjectName and StdoutUri are handled by msh to make
streaming between environments seamless.

Block comment for formatting and these are just notes.

```
echo hello | msh aws -- ./app.rb
    msh handles:
        * setting up aws environment configuration
        * executing ./app.rb
    app.rb also has '#!/usr/bin/env msh lambda ruby' which:
        * deploying shell script as a lambda ruby function or whatever
        * stdin lines are converted to lambda calls of the remote app with APIGateway context
        * streams output from lambda to stdout
msh aws -- ./app1.rb | msh aws -- ./app.rb
    * msh for app1 detects stdout is not the terminal
        * local pipe contains structured info with stdout info
        * lines to lambda calls 
        * client invocation with APIGateway context (doesn't actually use api gateway, just context structure passed)
          * Allows easy integration with APIGateway later if needed.
        * Current ARN passed to next command for IAM permissions if needed
lambda stdout to lambda: client invocation with http context
    same
```
    

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

Unfinished but demonstrate remote to remote and remote to local.
```bash
                    msh remote ./app                             |                                msh remote ./app  
deploy if needed then stdin lines to http then destroy stdin EOF |   reads config from stdin, ARN's of app and SQS output queue to configure this app
```

## Local Invocation Example
```bash
% cat ../msh-repo/conf/ruby/app.rb|head -n 1
#!/usr/local/msh/msh lambda ruby
%
% ./app.rb
  <runs in lib>
```
        

## State Machine

The above idea's could be used for quick iteration of ideas while something like what I describe [here](https://github.com/RyanJarv/coderun#update-1142020) could be used for long running data processing. I typed up a simple and incomplete spec for how shell to state machine compiling could work but unsure where this went, will add it if I find it.
