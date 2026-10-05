package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
)

var fs = NewOSFileSystem(".")
var files = fs.walkDir(".", &WalkDirOptions{
	SkipDirs:  fs.readIgnoreFile(),
	FilesOnly: true,
})

type OSFileSystem struct {
	Root string
}

type WalkDirOptions struct {
	SkipDirs  []string
	FilesOnly bool
}

func NewOSFileSystem(root string) *OSFileSystem {
	return &OSFileSystem{Root: root}
}

func (fs *OSFileSystem) root() string {
	return fs.Root
}

func (fs *OSFileSystem) resolve(path string) string {
	return filepath.Join(fs.Root, path)
}

func (fs *OSFileSystem) exists(path string) bool {
	_, err := os.Stat(fs.resolve(path))
	return err == nil
}

func (fs *OSFileSystem) readFile(path string) []byte {
	file, err := os.ReadFile(fs.resolve(path))
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func (fs *OSFileSystem) glob(pattern string) []string {
	matches, err := filepath.Glob(fs.resolve(pattern))
	if err != nil {
		fmt.Println(err)
	}
	return matches
}

func (fs *OSFileSystem) walkDir(path string, options *WalkDirOptions) []string {
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

func (fs *OSFileSystem) readIgnoreFile() []string {
	return strings.Split(string(fs.readFile(".yokipackignore")), "\n")
}
