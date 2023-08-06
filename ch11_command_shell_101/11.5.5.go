package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var ErrWildcardNoMatchError = errors.New("wild card no match")

func expandPath(dir, workDir string) string {
	// dir: *.go
	// IsAbs reports whether the path is absolute.
	if filepath.IsAbs(dir) {
		return dir
	}
	// Join joins any number of path elements into a single path, separating them with an OS specific Separator.
	// filepath.Join(workDir, dir): /Users/rayt/go-learning/ch11_command_shell_101/*.go
	return filepath.Join(workDir, dir)
}

func expandWildcard(arg, workDir string) ([]string, error) {
	if !strings.ContainsAny(arg, "*?[") {
		return []string{arg}, nil
	}
	files, err := filepath.Glob(expandPath(arg, workDir))
	if len(files) == 0 {
		return nil, ErrWildcardNoMatchError
	}
	return files, err
}

func main() {
	workDir := "/Users/rayt/go-learning/ch11_command_shell_101"
	arg := "*.go" // Example wildcard pattern

	files, err := expandWildcard(arg, workDir)
	if err != nil {
		if err == ErrWildcardNoMatchError {
			fmt.Println("No matches found for wildcard:", arg)
		} else {
			fmt.Println("Error:", err)
		}
		return
	}

	fmt.Println("Matched files:")
	for _, file := range files {
		fmt.Println(file)
	}
}
