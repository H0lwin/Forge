BINARY ?= forge
LDFLAGS ?= -X forge/cmd.version=1.0.0 -X forge/cmd.commit=$$(git rev-parse --short HEAD 2>/dev/null || echo dev) -X forge/cmd.date=$$(date -u +%Y-%m-%dT%H:%M:%SZ)

.PHONY: build test test-e2e lint release-snapshot

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./

test:
	go test ./...

test-e2e: build
	bash scripts/e2e_smoke.sh

lint:
	go test ./...

release-snapshot: build
	@echo "snapshot built at bin/$(BINARY)"
