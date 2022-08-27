[CmdletBinding()]
param(
    [Parameter()][switch] $NoCover
)

# https://github.com/wgross/fswatcher-engine-event

$Command = "go test ./..."
if ($VerbosePreference) {
    $Command = $Command + " -v"
}
if (-not $NoCover) {
    $Command = $Command + " -cover"
}
Write-Host $Command -ForegroundColor Green
Invoke-Expression $Command
