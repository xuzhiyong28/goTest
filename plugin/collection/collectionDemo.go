package collection

import (
	"fmt"
	"github.com/chenhg5/collection"
)


func Demo1(){
	a := []string{"h", "e", "l", "l", "o"}
	aa := collection.Collect(a).All()
	fmt.Println(aa)

	floats := []float64{143.66, -14.55}
	floats_s := collection.Collect(floats).Avg().String()
	fmt.Println(floats_s)

	array := collection.Collect(a).Chunk(3).ToMultiDimensionalArray()
	fmt.Println(array[0])

	nums := []int{4, 5, 2, 3, 6, 7}
	fmt.Println(collection.Collect(nums).Sort().ToInt64Array())

	//打乱数组
	fmt.Println(collection.Collect(nums).Shuffle().ToIntArray())

	fmt.Println("是否包含指定元素" , collection.Collect(nums).Contains(4))

	fmt.Println("长度" , collection.Collect(nums).Count())
}

func Demo2(){
	a := map[string]interface{}{
		"name": "mike",
		"sex":  1,
	}
	fmt.Println(collection.Collect(a).Keys().ToStringArray())
}