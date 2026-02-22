package domain

import (
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

var slugRe = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{1,62}$`)

func ValidateName(name string) error {
	if !slugRe.MatchString(name) {
		return fmt.Errorf("%w: --name must be lowercase slug (2-63 chars, a-z0-9-)", ErrValidation)
	}
	return nil
}

func ResolveProjectPath(basePath, name string) (string, error) {
	if strings.TrimSpace(basePath) == "" {
		return "", fmt.Errorf("%w: --path is required", ErrValidation)
	}
	path := filepath.Join(basePath, name)
	return filepath.Clean(path), nil
}

func RequiresPythonVersion(f Framework) bool {
	return f == FrameworkDjango || f == FrameworkFastAPI || f == FrameworkFlask
}

func RequiresEnvManager(f Framework) bool {
	return RequiresPythonVersion(f)
}

func ValidateFrameworkExtras(f Framework, extras []string) error {
	for _, ex := range extras {
		switch ex {
		case "drf":
			if f != FrameworkDjango {
				return fmt.Errorf("%w: extra %q is only supported for django", ErrValidation, ex)
			}
		case "celery":
			if !slices.Contains([]Framework{FrameworkDjango, FrameworkFastAPI, FrameworkFlask}, f) {
				return fmt.Errorf("%w: extra %q is only supported for python backends", ErrValidation, ex)
			}
		case "tailwind":
			if !slices.Contains([]Framework{FrameworkNext, FrameworkVite}, f) {
				return fmt.Errorf("%w: extra %q is only supported for frontend frameworks", ErrValidation, ex)
			}
		}
	}
	return nil
}

func MissingRequiredFlags(req GenerateRequest) []string {
	var missing []string
	if strings.TrimSpace(req.Name) == "" {
		missing = append(missing, "--name")
	}
	if strings.TrimSpace(string(req.Framework)) == "" {
		missing = append(missing, "--framework")
	}
	if strings.TrimSpace(req.BasePath) == "" {
		missing = append(missing, "--path")
	}
	if RequiresPythonVersion(req.Framework) && strings.TrimSpace(req.PythonVersion) == "" {
		missing = append(missing, "--python-version")
	}
	if RequiresEnvManager(req.Framework) && strings.TrimSpace(req.EnvManager) == "" {
		missing = append(missing, "--env-manager")
	}
	return missing
}
