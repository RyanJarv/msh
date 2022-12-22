# cmds="cat grep fgrep egrep awk bzip2 bzcat echo gzip nc pgrep ping ping6 sha1sum sha3sum sha256sum sha512sum sleep strings tar tac tee wget xz xzcat sort rev"

cmds="grep echo awk sort rev"

mkdir -p ./zip/bin

for cmd in $cmds; do
	curl -L -o "zip/bin/$cmd" "https://busybox.net/downloads/binaries/1.35.0-x86_64-linux-musl/busybox_$(echo $cmd|tr 'a-z' 'A-Z')"
	chmod +x "zip/bin/$cmd"
done
