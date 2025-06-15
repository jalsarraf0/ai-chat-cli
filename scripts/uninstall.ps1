Stop-Service -Name ai-chat-cli -ErrorAction SilentlyContinue
Remove-Item "$Env:ProgramFiles\ai-chat-cli" -Recurse -Force -ErrorAction SilentlyContinue
scoop uninstall ai-chat-cli   2>$null
Write-Host "✅ ai-chat-cli removed."
