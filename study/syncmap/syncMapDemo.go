package syncmap

import (
	"fmt"
	"sync"
)

/****
	传统模式下go的map结构在并发环境下会有并发问题，我们一般可以通过加锁的方式解决，但是这样很慢
	go给我们提供了并发安全的map结构 - sync.map  开箱即用
 */

func Demo1(){
	var scene sync.Map
	//保存值
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	//从sync.map中获取值
	fmt.Println(scene.Load("greece"))

	// 根据键删除对应的键值对
	scene.Delete("london")

	//遍历所有的值
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

	// 其他方法请参考api
}