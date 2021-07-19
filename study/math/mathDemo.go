package math

import (
	"fmt"
	"math/big"
)

func Demo1() {
	z := big.NewInt(0)
	a := big.NewInt(100)
	b := big.NewInt(11)
	z.Sub(a, b)
	fmt.Println("Sub = ", z)

	z.Add(a, b)
	fmt.Println("Add = ", z)

	z.Div(a, b)
	fmt.Println("Div = ", z)

	z.And(a, b) // 按位且
	fmt.Println("And = ", z)

	z.Or(a, b) //按位或
	fmt.Println("Or = " , z)

	z.Lsh(a,2) // a左位移运算
	fmt.Println("Lsh = ", z)

	j := big.NewInt(-10)
	fmt.Println(j.Sign()) //返回x的正负号。x<0时返回-1；x>0时返回+1；否则返回0。

}

/***
大数相加
 */
func Demo2(){
	a, _ := big.NewInt(0).SetString("9122322238215458478512545454878168716584545412154785452142499999", 10)
	fmt.Println(a)
	b := big.NewInt(2)
	fmt.Println(b)
	c := b.Add(a,b)
	fmt.Println(c)
}