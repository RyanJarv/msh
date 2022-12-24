export PATH_BAK=$PATH
unset PATH

function command_not_found_handler() {
  echo "Running: ./out/msh.pipes $@"
  ./out/msh.pipes "$@"
}
