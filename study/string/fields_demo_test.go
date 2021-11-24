package string

import (
	"fmt"
	"testing"
)

func TestDemo1(t *testing.T) {
	FieldsDemo()
}

func BenchmarkStringPlus10(b *testing.B) {
	p := initStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringPlus(p)
	}
}

func BenchmarkStringFmt10(b *testing.B) {
	p := initStringi(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringFmt(p)
	}
}

func TestOther2(t *testing.T) {
	var a int64 = 10
	fmt.Printf("变量a对应的地址 %p", &a)
	fmt.Println()
	modifiedNumber2(&a) // args就是实际参数
	fmt.Printf("最终a的值 %d", a)
	fmt.Println()
}

func modifiedNumber2(c *int64) { //这里定义的args就是形式参数
	fmt.Printf("变量c存储的值 %v" , c)
	fmt.Println()
	*c = 100
}

func TestOther1(t *testing.T) {
	var a int64 = 10
	fmt.Printf("变量a对应的地址 %p\n", &a)
	modifiedNumber(a)
	fmt.Printf("最终a的值 %d", a)
}

func modifiedNumber(c int64) {
	fmt.Printf("变量c对应的地址 %p\n", &c)
	c = 100
}
