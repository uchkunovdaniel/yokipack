package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
)

type OSFileSystem struct {
	root string
}

type WalkDirOptions struct {
	SkipDirs  []string
	FilesOnly bool
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

func (fs *OSFileSystem) WalkDir(path string, options *WalkDirOptions) []string {
	var dirs []string
	err := godirwalk.Walk(fs.resolve(path), &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if options.FilesOnly && de.IsDir() {
				return nil
			}
			if len(options.SkipDirs) > 0 {
				for _, s := range options.SkipDirs {
					if s != "" && strings.Contains(osPathname, s) {
						if b, err := de.IsDirOrSymlinkToDir(); b == true && err == nil {
							return filepath.SkipDir
						}
						return nil
					}
				}
			}
			dirs = append(dirs, osPathname)
			return nil
		}})
	if err != nil {
		fmt.Println(err)
	}
	return dirs
}
