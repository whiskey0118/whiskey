package common

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func GetPath(fileName string) string {
	paths := make([]string, 0)
	if execFile, err := os.Executable(); err == nil {
		paths = append(paths, filepath.Join(filepath.Dir(execFile), fileName))
	}
	if _, srcFile, _, ok := runtime.Caller(0); ok {
		paths = append(paths, filepath.Join(filepath.Dir(srcFile), fileName))
	}
	// Same folder of working folder
	if workingDir, err := os.Getwd(); err == nil {
		paths = append(paths, filepath.Join(workingDir, fileName))
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			log.Printf("using %v\n", p)
			return p
		}
	}
	return ""
}
