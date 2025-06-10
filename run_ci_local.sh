#!/usr/bin/env bash
set -Eeuo pipefail
LOG_DIR="logs"
IMAGE="ghcr.io/catthehacker/ubuntu:act-24.04"
WORKFLOW=".github/workflows/ci.yml"

timestamp() { date +"%Y%m%d-%H%M%S"; }
log_file="$LOG_DIR/ci_local_$(timestamp).log"

mkdir -p "$LOG_DIR"

echo "🛠  Checking for nektos/act …"
# Look for act in PATH or ./bin
if ! command -v act >/dev/null 2>&1; then
  if [[ -x ./bin/act ]]; then
    echo "   • found ./bin/act → adding to PATH"
    export PATH="$PWD/bin:$PATH"
  else
    echo "   • act not found → installing to /usr/local/bin"
    curl -sSL https://raw.githubusercontent.com/nektos/act/master/install.sh \
      | sudo bash -s -- -b /usr/local/bin
  fi
fi
echo "   • act $(act --version) installed"

echo "🐳  Pulling runner image $IMAGE …"
docker pull -q "$IMAGE"

SECRETS_FILE=".secrets"
if [[ ! -s $SECRETS_FILE ]]; then
  echo "⚠️  .secrets missing — creating placeholder"
  cat > "$SECRETS_FILE" <<'EOF'
GITHUB_TOKEN=
EOF
fi

echo "🚀  Running CI workflow locally"
echo "    ▶ logs: $log_file"
echo "------------------------------------------------------------------"

act --verbose \
    --workflows "$WORKFLOW" \
    --secret-file "$SECRETS_FILE" \
    --bind \
    -P ubuntu-latest="$IMAGE" \
    | tee "$log_file"

echo "✅  Finished — exit status ${PIPESTATUS[0]}"
