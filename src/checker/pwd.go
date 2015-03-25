package checker

import (
	"fmt"
	"os"
	"path/filepath"
)

func IsGitRepository() (b bool) {
	if path, err := os.Getwd(); err == nil {
		for {
			if isGitRepository(path) {
				return true
			} else {
				if parent := filepath.Dir(path); parent == path {
					break
				} else {
					path = parent
				}
			}
		}
	}
	return false
}

func isGitRepository(path string) (b bool) {
	dotgit_path := fmt.Sprintf("%s%c%s", path, os.PathSeparator, ".git")
	_, e := os.Stat(dotgit_path)
	return e == nil
}
