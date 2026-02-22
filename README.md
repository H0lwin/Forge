# forge

Production-grade scaffolding CLI written in Go.

## Build

```bash
go build -o bin/forge ./
```

Windows:

```powershell
go build -o bin/forge.exe ./
```

## Run

```bash
./bin/forge
./bin/forge new
./bin/forge doctor
```

## Non-interactive examples

```bash
./bin/forge new \
  --name my-api \
  --framework django \
  --path . \
  --python-version 3.11 \
  --env-manager venv \
  --extras docker,ci,drf,postgres \
  --no-interactive

./bin/forge new --name my-app --framework next --path . --extras docker --no-interactive --dry-run
```

## Test

```bash
go test ./...
```

## CI smoke scripts

```bash
bash scripts/e2e_smoke.sh
```

Windows:

```powershell
./scripts/e2e_smoke.ps1
```
