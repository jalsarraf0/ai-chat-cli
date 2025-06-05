function Install-Tools {
    $Env:PATH = "$(Get-Location)\offline-bin;$Env:PATH"
    New-Item -ItemType Directory -Force -Path ./offline-bin | Out-Null

    if (-not ((go version) -match 'go1\.24')) {
        Write-Error 'Go 1.24.x is required.'; exit 1
    }

    $missing = @('gofumpt','staticcheck','gosec') |
               Where-Object { -not (Get-Command $_ -ErrorAction SilentlyContinue) }
    if (-not $missing) { return }

    Write-Host "installing: $($missing -join ', ')" -ForegroundColor Cyan

    foreach ($pkg in @(
        'mvdan.cc/gofumpt@latest',
        'honnef.co/go/tools/cmd/staticcheck@latest',
        'github.com/securego/gosec/v2/cmd/gosec@latest'
    )) { go install $pkg --% -trimpath }
}
Install-Tools
