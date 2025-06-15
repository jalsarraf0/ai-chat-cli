$ErrorActionPreference = 'SilentlyContinue'

if (Get-Service -Name 'ai-chat' -ErrorAction SilentlyContinue) {
    Stop-Service 'ai-chat'
    sc.exe delete 'ai-chat' | Out-Null
}

$bins = @(
    (Get-Command ai-chat -ErrorAction SilentlyContinue).Source,
    "$env:USERPROFILE\scoop\apps\ai-chat-cli\current\ai-chat.exe"
)
foreach ($b in $bins) {
    if ($b -and (Test-Path $b)) { Remove-Item $b -Force }
}

$config = Join-Path (${env:XDG_CONFIG_HOME} ? $env:XDG_CONFIG_HOME : "$env:APPDATA") 'ai-chat-cli'
Remove-Item $config -Recurse -Force -ErrorAction SilentlyContinue

if (Get-Command scoop -ErrorAction SilentlyContinue) { scoop uninstall ai-chat-cli }
if (Get-Command brew -ErrorAction SilentlyContinue) { brew uninstall --force ai-chat-cli }

Write-Host 'ai-chat-cli uninstalled'
