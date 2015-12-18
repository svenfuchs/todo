package test

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func ExpandRelativePath(path string) string {
	_, file, _, _ := runtime.Caller(1)
	return filepath.Clean(file + "/" + path)
}

func ReadFile(path string) (string, error) {
	str, err := ioutil.ReadFile(path)
	return string(str), err
}
