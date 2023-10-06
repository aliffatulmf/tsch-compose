$env:GOOS = "windows"
$env:GOARCH = "amd64"

Write-Output "Building from source"
Remove-Item -Path "tsch-debug" -Recurse -Force -ErrorAction Ignore 
New-Item -ItemType Directory -Path "tsch-debug"

go build -o "tsch-debug/tsch-compose.exe" main.go

Copy-Item -Path "$PWD\example\tsch-compose.yaml" -Destination "$PWD\tsch-debug\tsch-compose.yaml" -Force

Write-Output "Running tests"
Set-Location -Path "tsch-debug"

.\tsch-compose.exe

try {
    $result = SCHTASKS /QUERY /FO TABLE | Select-String -Pattern "Task 1.1"
    if ($null -eq $result) {
        throw "Task not found"
    }
}catch {
    Write-Error "$_"
}