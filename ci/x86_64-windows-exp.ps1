$env:GOOS = "windows"
$env:GOARCH = "amd64"

Write-Host "Building from source"
Remove-Item -Path "tsch-debug" -Recurse -Force -ErrorAction Ignore 
New-Item -ItemType Directory -Path "tsch-debug"

go build -o "tsch-debug/tsch-compose.exe" main.go

Copy-Item -Path "$PWD\example\tsch-compose.yaml" -Destination "$PWD\tsch-debug\tsch-compose.yaml" -Force

Write-Host "Running tests"
Set-Location -Path "tsch-debug"

.\tsch-compose.exe up --verbose

Write-Host "Checking if task is created"

try {
    $task = Get-ScheduledTask -TaskName "Task Daily" -ErrorAction Stop
    if ($null -eq $task) {
        throw "Task not found"
    }
} catch {
    Write-Error "$_"
}
