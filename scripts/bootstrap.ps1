<#
  bootstrap.ps1 — Windows flavour of tool pre-flight
#>

function Install-Tools {
    param(
        [string[]]$Pkgs = @(
            'mvdan.cc/gofumpt@latest',
            'honnef.co/go/tools/cmd/staticcheck@latest',
            'github.com/securego/gosec/v2/cmd/gosec@latest',
            'golang.org/x/vuln/cmd/govulncheck@latest',
            'github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.1'
        )
    )

    $Env:PATH = "$(Get-Location)\offline-bin;$Env:PATH"
    New-Item -ItemType Directory -Force -Path ./offline-bin | Out-Null

    if (-not ((go version) -match 'go1\.24')) {
        Write-Error 'Go 1.24.x is required.'; exit 1
    }

    Write-Host '🔧 installing tools' -ForegroundColor Cyan

    foreach ($pkg in $Pkgs) {
        go install -trimpath $pkg
    }
}

Install-Tools

