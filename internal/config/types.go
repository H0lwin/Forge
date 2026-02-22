package config

type Config struct {
	SchemaVersion int `mapstructure:"schema_version" yaml:"schema_version"`
	User          struct {
		Name   string `mapstructure:"name" yaml:"name"`
		Email  string `mapstructure:"email" yaml:"email"`
		Github string `mapstructure:"github" yaml:"github"`
	} `mapstructure:"user" yaml:"user"`
	Defaults struct {
		GitInit    bool   `mapstructure:"git_init" yaml:"git_init"`
		OpenEditor bool   `mapstructure:"open_editor" yaml:"open_editor"`
		Editor     string `mapstructure:"editor" yaml:"editor"`
	} `mapstructure:"defaults" yaml:"defaults"`
	Python struct {
		DefaultVersion string `mapstructure:"default_version" yaml:"default_version"`
		EnvManager     string `mapstructure:"env_manager" yaml:"env_manager"`
	} `mapstructure:"python" yaml:"python"`
	Node struct {
		PackageManager string `mapstructure:"package_manager" yaml:"package_manager"`
	} `mapstructure:"node" yaml:"node"`
}

func Default() Config {
	cfg := Config{SchemaVersion: 1}
	cfg.Defaults.GitInit = true
	cfg.Defaults.OpenEditor = false
	cfg.Defaults.Editor = "code"
	cfg.Python.DefaultVersion = "3.11"
	cfg.Python.EnvManager = "venv"
	cfg.Node.PackageManager = "pnpm"
	return cfg
}
