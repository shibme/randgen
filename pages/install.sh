#!/bin/sh

set -e

if ! command -v unzip >/dev/null && ! command -v 7z >/dev/null; then
	echo "Error: either unzip or 7z is required to install RandGen" 1>&2
	exit 1
fi

if [ "$OS" = "Windows_NT" ]; then
	target="windows-amd64"
else
	case $(uname -sm) in
	"Darwin x86_64") target="darwin_amd64" ;;
	"Darwin arm64") target="darwin_arm64" ;;
	"Linux aarch64") target="linux_arm64" ;;
	*) target="linux_amd64" ;;
	esac
fi

if [ $# -eq 0 ]; then
	randgen_uri="https://github.com/shibme/randgen/releases/latest/download/randgen_${target}.zip"
else
	randgen_uri="https://github.com/shibme/randgen/releases/download/${1}/randgen_${target}.zip"
fi

randgen_install="${RANDGEN_INSTALL:-$HOME/.randgen}"
bin_dir="$randgen_install/bin"
exe="$bin_dir/randgen"

if [ ! -d "$bin_dir" ]; then
	mkdir -p "$bin_dir"
fi

curl --fail --location --progress-bar --output "$exe.zip" "$randgen_uri"
if command -v unzip >/dev/null; then
	unzip -d "$bin_dir" -o "$exe.zip"
else
	7z x -o"$bin_dir" -y "$exe.zip"
fi
chmod +x "$exe"
rm "$exe.zip"

echo "RandGen was installed successfully to $exe"
if command -v randgen >/dev/null; then
	echo "Run 'randgen --help' to get started"
else
	case $SHELL in
	/bin/zsh) shell_profile=".zshrc" ;;
	*) shell_profile=".bashrc" ;;
	esac
	echo "Manually add the directory to your \$HOME/$shell_profile (or similar)"
	echo "  export RANDGEN_INSTALL=\"$randgen_install\""
	echo "  export PATH=\"\$RANDGEN_INSTALL/bin:\$PATH\""
	echo "Run '$exe --help' to get started"
fi
echo
