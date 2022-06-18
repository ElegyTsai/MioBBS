package utils

import (
	"runtime"
	"strings"
)

func GetProjectPath() string {
	_, filename, _, _ := runtime.Caller(1)
	//return filename
	return string([]rune(filename)[:3+strings.LastIndex(filename, "src")])
}
