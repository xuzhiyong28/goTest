package panic

import "fmt"

func panicDemo1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误", err)
		}
	}()

	fmt.Println("a")
	panic("n")
	fmt.Println("b") //这里不会在执行了
}


func panicDemo2(){
	fn1()
	fn2() //此方法报错，但不影响后面的执行，因为做了recover处理
	fmt.Println("panicDemo2")
}

func fn1(){
	fmt.Println("fn1")
}

func fn2(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("错误", err)
		}
	}()
	panic("fn2")
}