package engine

import (
	"fmt"
	"os"
	"path/filepath"
)

type OSFileSystem struct {
	root string
}

func NewOSFileSystem(root string) *OSFileSystem {
	return &OSFileSystem{root: root}
}

func (fs *OSFileSystem) Root() string {
	return fs.root
}

func (fs *OSFileSystem) resolve(path string) string {
	return filepath.Join(fs.root, path)
}

func (fs *OSFileSystem) Exists(path string) bool {
	_, err := os.Stat(fs.resolve(path))
	return err == nil
}

func (fs *OSFileSystem) ReadFile(path string) []byte {
	file, err := os.ReadFile(fs.resolve(path))
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func (fs *OSFileSystem) Glob(pattern string) []string {
	matches, err := filepath.Glob(fs.resolve(pattern))
	if err != nil {
		fmt.Println(err)
	}
	return matches
}

func (fs *OSFileSystem) ReadDir(path string) []os.DirEntry {
	dir, err := os.ReadDir(fs.resolve(path))
	if err != nil {
		fmt.Println(err)
	}
	return dir
}
