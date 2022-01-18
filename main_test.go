package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

// 整个文件读取入内存 适合小文件，大文件不适合
func TestFile1(t *testing.T) {
	// 方式1 ：使用os
	content, err := os.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	// 方式2：使用ioutil
	content2, err2 := ioutil.ReadFile("a.txt")
	if err2 != nil {
		panic(err)
	}
	fmt.Println(string(content2))

	// 方式3：使用os.open + ioutil 采用先创建句柄在读取
	file, err := os.Open("a.txt") //只读的方式，所以更安全
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content3, _ := ioutil.ReadAll(file)
	fmt.Println(string(content3))
}

// 每次读取一行 适用于大文件，不用占用大量内存
func TestFile2(t *testing.T) {
	// 方式1 ： bufio.ReadBytes('\n') 通过\n分隔符读取
	fi, _ := os.Open("christmas_apple.py")
	r := bufio.NewReader(fi)
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}

	// 方式2 ：bufio.ReadString('\n') 通过\n分隔符读取
	r2 := bufio.NewReader(fi)
	for {
		line, err := r2.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(line)
	}
}

// 每次只读取固定字节数 适用于大文件
func TestFile3(t *testing.T) {
	fi, _ := os.Open("a.txt")
	r := bufio.NewReader(fi)
	// 每次读取 1024 个字节
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
