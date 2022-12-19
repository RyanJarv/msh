# Use a fifo queue?

* For most stuff only closed events have to be processed after all other messages.



# In order commands
  * sort can operate on an out-of-order file that contains all messages.
  * uniq needs to be in order or ran on the output of sort.
    * Might be able to delay execution and run these as one step.
  * shutdown events, see below

Order needs to be tracked through out-of-order steps, so it can be restored when needed. Whether order is needed
may also depend on command line flags used, for example, `grep -A 10`.

## Shutdown events

We can't signal when shutdown should happen for the next command if no state exists across lambda startups. We
need to know what has been processed at the current step for this to work. Lambda state is not shared or saved, 
however.

### Sharing State

* Store processed count in dynamodb for each command.
    * Believe you can do atomic increments.
    * On shutdown, we send a new closed event to the next command with the count of records sent.
    * SQS may have something built in we could use here?
* Communication between remote and local processes.
    * Send processed count over SQS, which is read and counted locally.
    * Pretty similar to dynamodb.

### When state is shared

State doesn't need to be shared at all steps, for shutdown we can send noops for filtered events which get collected
at the end of the pipeline. This would also allow us to detect any accidentally dropped messages as well as know when
all messages in the out-of-order segment have been received.

### Signalling shutdown backwards through the pipeline.

In a unix shell, is signaled the other way, stdin of the next command is closed indicating the end of the input. To
avoid sharing state at each step, we postpone deciding when to shut down in the local pipeline until a collection is
reached (for now, one that aggregates output to the tty).

Unix pipes work one way, in `first | second`, so `second` has no pipe to send data to `first`.

Options:
  * Aggregate pipeline ARNs when forwarding the config during setup, the collection command then closes these at 
    shutdown.
  * Each local command manages its own shutdown.
    * `first` can detect when `second` closes stdin by attempting to send to it.
        * All commands, even local ones would need to support this. This may be ok as we need a wrapper for non-msh commands
          to work anyway because events would need to be collected from SQS before sending to the subprocess. We should also
          be able to configure a catch-all wrapper in the shell to avoid supporting all commands, same way ubuntu does this
          with the command not found thing (below).
    * Only send the local PID of the first msh command in the pipeline. The collection command SIGTERM's this process at
      shutdown, which shuts down the rest of the pipeline normally.
      * The startup PID will be the one where stdin is a tty or doesn't receive a msh startup event from the previous 
        command.
      * This behavior is a bit closer to how normal shell pipelines work, reading from a closed stdin causes the 
        current command to exit.

Signals, like ctrl-c go to all processes, so this should behave fairly similar regardless.

## Shutdown

Commands that have stdin set to a tty send a closed event locally with the number of messages sent.

```
$ first | cat
{"Event": "Shutdown", "Processed": 5}
```

```
$ first | second | cat
{"Event": "Shutdown", "Processed": 5}
```
This is forwarded through as is.

In the remote processes we forward the number of messages received, if we only forward 5/10 we still indicate we 
received 10 in at least one message.

### Noop filtering

input: 
> [{"body": ...}, <2 more msgs>, {"filtered": 2}]

Filtered event indicates two where dropped in previous steps.

Lambda wrapper:
* Sum total of received filtered events.
* Pass non filter lines to sub cmd.
* Track the difference between input and output lines.
* If diff > 0; update filtered total.
* Add summed filter event to batch response.

output:
> [{"body": ...}, {"filtered": 4}]
 
Two more where dropped at this step, send the new total along.

### Collection command

Locally, we count the returned messages, adding counts from all filtered events. We know we have processed all 
messages when this equals the number sent to us in the shutdown event.

The PID of the first msh command, sent during the local config event, is sent a signal to indicate it should clean 
up and close stdout. This terminates the rest of the pipeline similar to any other non-msh cmds.

### Clean up

For now probably just stopping the eventbridge pipe, processing should stop when we think it's stopped.

## FIFO Queue

This doesn't seem to be working with eventbridge pipes.

Maybe related to the deduplication id? Unsure how this is supposed to be configured exactly.

# Catch-All Wrapper

## Zsh
% export PATH_BAK=$PATH
% unset PATH
% ls
zsh: command not found: ls
% command_not_found_handler() { echo "running $@"; PATH=$PATH_BAK "${=@}"; }
% ls
running ls
Dockerfile.base	NOTES.md	TODO.md		go.mod		main.py		pkg		test
Makefile	README.md	cmd		go.sum		out		scripts		utils
should

## Bash

Should work the same. Although older versions of bash, like the default version on macos, don't seem to support 
command_not_found_handler.

# Scratch

* Right now the last command pipes output back locally, so we could have the last command track this.
  ~~* Would break if the assumption is broken.
    * We could signal shutdown by closing stdin (stdout of the last command)...
        * A non-msh cat running as the first command would not exit here. (`cat | sleep 1` vs `sleep 1 | cat`)
            * I think this is because most commands will block on stdin, not normal to be watching stdout when no data
              is being processed.
            * Could have a catch-all wrapper. May need this anyways?~~
    * This won't work
        *
      across separate lambda startups. We can't signal what should be delivered without knowing when
      * Need to track state across lambda executions to know how many have been delivered.
        * Timestamps of last delivered message don't make a difference here.
            * Still n
