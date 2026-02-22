package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func Validate(cfg Config) error {
	if cfg.SchemaVersion <= 0 {
		return fmt.Errorf("schema_version must be positive")
	}
	if cfg.Python.EnvManager == "" {
		return fmt.Errorf("python.env_manager cannot be empty")
	}
	if cfg.Node.PackageManager == "" {
		return fmt.Errorf("node.package_manager cannot be empty")
	}
	return nil
}

func DefaultPath() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home dir: %w", err)
	}
	return filepath.Join(h, ".forge", "config.yaml"), nil
}

func Load(path string) (Config, error) {
	cfg := Default()
	if path == "" {
		p, err := DefaultPath()
		if err != nil {
			return cfg, err
		}
		path = p
	}
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("decode config: %w", err)
	}
	if err := Validate(cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func Save(path string, cfg Config) error {
	if err := Validate(cfg); err != nil {
		return err
	}
	if path == "" {
		p, err := DefaultPath()
		if err != nil {
			return err
		}
		path = p
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
