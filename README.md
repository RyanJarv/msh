## Multivac Shell

Currently this is just a PoC, it can not act as a sandbox currently (and may even allow escalation to root or other users on your host from the container as is).

### Goal

A lightweight, easy to use, sandboxed (WIP), userland.

### Example

```
% type python.msh
python.msh is /Users/me/.msh/bin/python.msh
% cat /Users/me/.msh/bin/python.msh
#!/usr/bin/env msh
FROM python:3
WORKDIR /app
ENTRYPOINT ["python"]
% python.msh
Python 3.10.5 (main, Jun 23 2022, 10:42:52) [GCC 10.2.1 20210110] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> print("hello from a container!")
hello from a container!
>>>
```

## Things To Know

* This project is a work in progress.
* Currently it is not safe to run untrusted code with msh.
* It shouldn't be used currently for anything other then a PoC.
* At the moment the current working directory is always shared with the container.

#### Features

* Container builds are done on the fly when a msh file is first executed.
* When the source msh file is modified, builds will run on next execution.

#### Future Features

* Support more complex dockerfiles.
  * Currently, builds are run an empty temp directory.
* Additional runtime options for sharing.
* Running on a remote system.
  * Previously this supported ECS and lambda but I have since removed that to simplify things for the moment.

### Install
```
git clone https://github.com/RyanJarv/msh.git
cd msh
make install
source ~/.msh/env
echo 'source ~/.msh/env' | tee -a ~/.bashrc ~/.zshrc
```


### Create your own msh file

Create a normal Dockerfile, make sure the working directory is /app and add a msh hash bang (`#!...`) to the first line.

```
$ cat ./python.msh
#!/usr/bin/env msh
FROM python:3
WORKDIR /app
ENTRYPOINT ["python"]
```

Ensure the file is executable.

```shell
chmod +x ./python.msh
```

You can now run the file as if it was a normal binary.

```
$ python.msh ./main.py 
hello from docker!
```

To run it anywhere, ensure it is moved to a directory in your path.

```shell
cp python.msh ~/.msh/bin/python.msh
echo 'source ~/.msh/env' | tee -a ~/.bashrc ~/.zshrc
```
