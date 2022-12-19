## Easy TODO

* Match owner permissions

## Thoughts

### Config

I think getting configuration format right is important, let's focus on that now.

### Platforms

Running locally was only one part of the goal of this project.

So far I've messed around with:

* lambda
  * limited execution time
* fargate
  * start up time is very slow
* kubernetes
  * too large of an upfront investment for the user
    * aws
      * 70 a month

Realistically whatever I use is likely going to effect the design of the project, so may be worth thinking about this
more.

Possible options:
  * gcp autopilot
    * .10 an hour it seems
