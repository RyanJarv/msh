function command_not_found_handler() {
  if [[ -z $MSH_ENABLED ]]; then
    echo "Command not found: $1"
    return
  fi

  ./out/msh.pipes "$@"
}

function msh() {
  export PATH_BAK="$PATH"
  export MSH_ENABLED=1

  unset PATH

  "$@"
}

function restore_msh() {
  if [[ -n $MSH_ENABLED ]]; then
    unset MSH_ENABLED
    export PATH="$PATH_BAK"
  fi
}

alias 'p{'='{ msh; '

# This doesn't always restore the path correctly.
alias -g '}'='; restore_msh; }; '