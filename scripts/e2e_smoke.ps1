$ErrorActionPreference = 'Stop'
$bin = if ($env:BIN) { $env:BIN } else { '.\\bin\\forge.exe' }
$tmp = New-Item -ItemType Directory -Path (Join-Path $env:TEMP ("forge-e2e-" + [guid]::NewGuid().ToString()))

& $bin version
& $bin doctor --no-interactive
& $bin new --framework django --name e2e-dj --path $tmp.FullName --python-version 3.11 --env-manager venv --extras drf,docker,ci --no-interactive --dry-run
& $bin new --framework next --name e2e-next --path $tmp.FullName --extras tailwind --no-interactive --dry-run
Write-Host "E2E smoke complete"
