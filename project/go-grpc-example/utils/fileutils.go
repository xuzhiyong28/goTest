package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetFileRealPath(src string) string {
	pwd, _ := os.Getwd()
	osPart := "/"
	if runtime.GOOS == "windows" {
		osPart = "\\"
	}
	src = pwd + osPart + src
	return src
}
