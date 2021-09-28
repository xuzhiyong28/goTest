package main

import "example/plugin/cobra/cmd"

/**
	1. go build -o main.exe
    2. ./main.exe -h   ./main.exe version
*/

func main() {
	cmd.Execute()
}