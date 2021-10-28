package main

import (
	"fmt"
	"os"
)

func main() {
	osArgs()
}

// go run main.go 1 2 3 4 5 6
func osArgs() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
