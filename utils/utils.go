package utils

import (
	"path/filepath"
	"runtime"
)

func GetInputPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Join(filepath.Dir(filename), "input.txt")
}
