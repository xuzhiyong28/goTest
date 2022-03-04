package _interface

import (
	"fmt"
	"testing"
)


func TestDemo0(t *testing.T){
	var greeting = "hello world"
	greetingStr , ok := interface{}(greeting).(string) // 解释下意思，类型断言只能作用在interface类型上，所以要先将greeting转成interface{}
	fmt.Println(fmt.Sprintf("greetingStr : %v, ok : %v", greetingStr, ok))
}


// 类型断言
// 类型断言用于断言变量是属于某种类型 , 类型断言只能发生在interface{}类型上。
func TestDemo1(t *testing.T) {
	var greeting interface{} = "hello world"
	greetingStr, ok := greeting.(string)
	fmt.Println(fmt.Sprintf("greetingStr : %v, ok : %v", greetingStr, ok))
}

// 类型转换
func TestDemo2(t *testing.T) {
	greeting := []byte("hello world")
	greetingStr := string(greeting)
	fmt.Println(fmt.Sprintf("greetingStr : %v" , greetingStr))
}