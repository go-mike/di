[CmdletBinding()]
param(
    [Parameter()][switch] $Cover
)

# https://github.com/wgross/fswatcher-engine-event

$Command = "go test ./..."
if ($VerbosePreference) {
    $Command = $Command + " -v"
}
if ($Cover) {
    $Command = $Command + " -covermode=count -coverprofile coverage.out"
} else {
    $Command = $Command + " -covermode=count"
}
Write-Host $Command -ForegroundColor Green
Invoke-Expression $Command

if ($Cover) {
    $Command = "go tool cover '-html=coverage.out'"
    Write-Host $Command -ForegroundColor Green
    Invoke-Expression $Command
}
