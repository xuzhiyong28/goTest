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

func TestOther(t *testing.T) {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[k] = 100 + v
	}
	fmt.Println(a)

	b := []int{1, 2, 3}
	for k, v := range b {
		if k == 0 {
			b[0], b[1] = 100, 200
			fmt.Println(b)
		}
		b[k] = 100 + v
	}
	fmt.Print(b)
}
