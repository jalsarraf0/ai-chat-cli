#!/usr/bin/env sh
set -e
sudo systemctl --quiet disable --now ai-chat-cli.service || true
sudo rm -f /usr/local/bin/ai-chat-cli /usr/bin/ai-chat-cli
sudo rm -rf /etc/ai-chat-cli /var/{lib,log}/ai-chat-cli
rpm -e ai-chat-cli 2>/dev/null || true
dpkg -r ai-chat-cli 2>/dev/null || true
command -v brew  && brew uninstall ai-chat-cli   || true
command -v scoop && scoop uninstall ai-chat-cli   || true
echo "✅ ai-chat-cli removed."
