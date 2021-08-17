package main

import (
	"fmt"
	"testing"
)

func TestGroup_Get(t *testing.T) {
	var slice []int
	slice = append(slice, 1, 2, 3)
	newSlice := AddElement(slice, 4)
	fmt.Println(&slice[0] == &newSlice[0]) //false
}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}
