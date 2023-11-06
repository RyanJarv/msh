export PS1="(msh) $PS1"
export PATH="$PWD/out:$PATH"

rehash || :

# TODO: Support the following:
#
#   @event events/patterns/ec2-running.json | sns | .{ …; … }.
#
# So that sns is connected to each topic in .{ ... }.
#
# Ideas 1:
#   alias '.{'='{ cat; '
#   alias ‘}.’=‘}| each;
#
#   echo a | .{ echo b; }. | cat
#   a
#   b
#
# Ideas 2:
#   alias '<('='{ cat; '
#   alias ‘}.’=‘}| each;
#
#   echo a | .{ echo b; }. | cat
#   a
#   b

function sfn {
  if [ $# -eq 0 ]; then
    command .sfn
  elif type "$name" | grep -q 'shell function'; then
    command .sfn "$name" | "$name" "$@"
  fi
}

# TODO: Come back to this
 alias 'sfn{'='.sfn | {'