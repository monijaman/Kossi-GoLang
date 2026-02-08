# Product Translation Tool Runner - Windows PowerShell

# This script makes it easy to run the product translator on Windows

# Change to script directory and then up to server root
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$serverRoot = Join-Path $scriptDir "..."

Push-Location $serverRoot

# Check if .env exists
if (-not (Test-Path ".env")) {
    Write-Host "❌ Error: .env file not found" -ForegroundColor Red
    Write-Host "Please create .env with DATABASE_URL" -ForegroundColor Yellow
    Pop-Location
    exit 1
}

# Parse command line arguments
$analyzeOnly = $false
$generateBN = $false
$showAll = $false
$verbose = $false

foreach ($arg in $args) {
    switch ($arg) {
        "analyze" { $analyzeOnly = $true }
        "generate" { $generateBN = $true }
        "all" { $showAll = $true }
        "verbose" { $verbose = $true }
    }
}

Write-Host "🚀 Starting Product Translator..." -ForegroundColor Green
Write-Host ""

# Build the flags
$flags = @()
if ($analyzeOnly) { $flags += "-analyze" }
if ($generateBN) { $flags += "-generate-bn" }
if ($showAll) { $flags += "-show-all" }
if ($verbose) { $flags += "-verbose" }

# Run the Go program
$command = "go run cmd/product-translator/main.go"
if ($flags.Count -gt 0) {
    $command += " " + ($flags -join " ")
}

Invoke-Expression $command

Pop-Location
