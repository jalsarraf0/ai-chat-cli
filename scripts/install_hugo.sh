#!/usr/bin/env bash
set -euo pipefail

version="0.124.1"
install_dir="$HOME/.local/bin"
mkdir -p "$install_dir"

real="$install_dir/hugo.real"
if [[ ! -x "$real" ]]; then
  url="https://github.com/gohugoio/hugo/releases/download/v${version}/hugo_extended_${version}_Linux-64bit.tar.gz"
  echo "🔧 Installing Hugo Extended ${version}"
  curl -sSL "$url" | tar -xz -C "$install_dir" hugo
  mv "$install_dir/hugo" "$real"
fi

wrapper="$install_dir/hugo"
cat > "$wrapper" <<'WRAP'
#!/usr/bin/env bash
args=()
for a in "$@"; do
  [[ "$a" == "--renderToREADME" ]] && continue
  args+=("$a")
done
exec "$(dirname "$0")/hugo.real" "${args[@]}"
WRAP
chmod +x "$wrapper"

export PATH="$install_dir:$PATH"
