#!/usr/bin/env sh
set -e

# stop service if running
if command -v systemctl >/dev/null 2>&1; then
    if systemctl list-units --full -all | grep -q ai-chat.service; then
        sudo systemctl disable --now ai-chat.service || true
        sudo rm -f /etc/systemd/system/ai-chat.service
        sudo systemctl daemon-reload
    fi
fi

# remove binaries
bins="/usr/local/bin/ai-chat /usr/bin/ai-chat $(go env GOPATH)/bin/ai-chat"
for b in $bins; do
    [ -e "$b" ] && sudo rm -f "$b" || true
done

# remove config and logs
cfg="${XDG_CONFIG_HOME:-$HOME/.config}/ai-chat-cli"
sudo rm -rf "$cfg" 2>/dev/null || true
sudo rm -rf /var/log/ai-chat-cli 2>/dev/null || true

# remove packages
if command -v dpkg >/dev/null 2>&1; then
    sudo dpkg -r ai-chat-cli 2>/dev/null || true
fi
if command -v rpm >/dev/null 2>&1; then
    sudo rpm -e ai-chat-cli 2>/dev/null || true
fi
if command -v brew >/dev/null 2>&1; then
    brew uninstall --force ai-chat-cli 2>/dev/null || true
fi
if command -v scoop >/dev/null 2>&1; then
    scoop uninstall ai-chat-cli 2>/dev/null || true
fi

echo "ai-chat-cli uninstalled"
