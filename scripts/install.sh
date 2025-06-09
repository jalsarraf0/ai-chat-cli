#!/usr/bin/env sh
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

DRY=0
YES=0
OPENAI=""
for arg in "$@"; do
    case "$arg" in
        --dry-run) DRY=1 ;;
        --yes) YES=1 ;;
        --openai-key=*) OPENAI="${arg#*=}" ;;
    esac
done
check(){ command -v "$1" >/dev/null 2>&1 || { echo "$1 missing"; exit 1; }; }
check go
check curl
check grep
check awk
check tar
if [ "$(uname)" = "Darwin" ]; then
    command -v brew >/dev/null 2>&1 || { echo "brew missing"; exit 1; }
fi
printf 'Welcome to ai-chat-cli installer\n'
[ $DRY -eq 1 ] && echo '(dry run)'
prompt(){
    var=$1; def=$2; msg=$3;
    [ $YES -eq 1 ] && { eval "$var=$def"; return; }
    printf '%s' "$msg"; read -r ans; [ -z "$ans" ] && ans=$def; eval "$var=$ans"
}
prompt MODEL "gpt-4o" "Model [gpt-4o]: "
prompt FORMAT "markdown" "Response format [markdown]: "
prompt TELEMETRY "no" "Enable telemetry? [y/N]: "
[ -n "$OPENAI" ] || prompt OPENAI "" "OPENAI_API_KEY: "
CONFIG_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/ai-chat-cli"
CONFIG="$CONFIG_DIR/config.yaml"
ENV_FILE="$CONFIG_DIR/.env"
do_install(){
    mkdir -p "$CONFIG_DIR"
    cat >"$CONFIG" <<EOF
model: $MODEL
format: $FORMAT
telemetry: $TELEMETRY
EOF
    chmod 600 "$CONFIG"
    if [ "$OPENAI" ]; then
        echo "OPENAI_API_KEY=$OPENAI" >"$ENV_FILE"
        chmod 600 "$ENV_FILE"
    fi
    (
        cd "$REPO_ROOT"
        go install ./cmd/ai-chat-cli@latest
    )
}
if [ $DRY -eq 1 ]; then
    echo "Would install to $CONFIG and build binary"
    exit 0
fi
do_install
ai-chat-cli --version
ai-chat-cli healthcheck
echo "Installation complete at $(go env GOBIN 2>/dev/null || echo "$HOME/go/bin")"
