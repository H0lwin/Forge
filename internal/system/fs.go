package system

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FileSystem interface {
	MkdirAll(path string, perm fs.FileMode) error
	WriteFile(path string, data []byte, perm fs.FileMode) error
	Stat(path string) (os.FileInfo, error)
	Exists(path string) bool
}

type OSFileSystem struct{}

func (OSFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (OSFileSystem) WriteFile(path string, data []byte, perm fs.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir parent: %w", err)
	}
	return os.WriteFile(path, data, perm)
}

func (OSFileSystem) Stat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

func (OSFileSystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
