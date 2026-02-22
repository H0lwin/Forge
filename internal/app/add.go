package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DetectFrameworkFromPath(path string) string {
	if path == "" {
		path = "."
	}
	checks := map[string][]string{
		"django":  {"manage.py", filepath.Join("config", "settings", "base.py")},
		"fastapi": {filepath.Join("app", "main.py")},
		"flask":   {"run.py", filepath.Join("app", "__init__.py")},
		"express": {filepath.Join("src", "index.js")},
		"nestjs":  {filepath.Join("src", "main.ts")},
		"next":    {filepath.Join("app", "page.tsx")},
		"vite":    {"index.html", filepath.Join("src", "main.tsx")},
	}
	for framework, files := range checks {
		ok := true
		for _, f := range files {
			if _, err := os.Stat(filepath.Join(path, f)); err != nil {
				ok = false
				break
			}
		}
		if ok {
			return framework
		}
	}
	return ""
}

func AddonFile(addon string) (string, string, error) {
	addon = strings.ToLower(addon)
	switch addon {
	case "auth":
		return "AUTH.md", "JWT auth setup\n\n1. Install dependency\n2. Configure middleware\n3. Protect routes/views\n", nil
	case "celery":
		return "CELERY.md", "Celery + Redis setup\n\n1. Add broker URL\n2. Create celery app module\n3. Run worker process\n", nil
	case "cache":
		return "CACHE.md", "Redis cache setup\n\nSet CACHE_URL and wire cache backend in your framework settings.\n", nil
	case "email":
		return "EMAIL.md", "Email backend setup\n\nConfigure SMTP or provider API key in .env and settings.\n", nil
	case "docker":
		return "Dockerfile", "FROM alpine:3.20\nWORKDIR /app\nCOPY . .\nCMD [\"sh\", \"-c\", \"echo run app command here\"]\n", nil
	case "ci":
		return filepath.Join(".github", "workflows", "ci.yml"), "name: ci\non: [push, pull_request]\njobs:\n  test:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n", nil
	case "sentry":
		return "SENTRY.md", "Sentry setup\n\nSet SENTRY_DSN in .env and initialize SDK in app startup.\n", nil
	case "pytest":
		return "pytest.ini", "[pytest]\naddopts = -q\n", nil
	default:
		return "", "", fmt.Errorf("unsupported addon: %s", addon)
	}
}
