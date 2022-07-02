## Multivac Shell

A lightweight, easy to use, sandboxed, userland.

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
* The use of the word "sandboxed" is relative here. You can probably think of this as something along the lines of user seperation, when possible, per executable.

#### Features

* Container builds are done on the fly when a msh file is first executed.
* When the source msh file is modified, builds will run on next execution.

#### Potential Future Features

* Support more complex dockerfiles.
  * Currently, builds are run an empty temp directory.

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
