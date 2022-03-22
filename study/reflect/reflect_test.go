//教程 ： https://www.liwenzhou.com/posts/Go/13_reflect/
package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type myInt int64

type person struct {
	name string
	age  int64
}

func (s person) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func (s person) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func TestReflect_demo0(t *testing.T) {
	var a float32 = 3.14
	reflectType(a)

	var b int64 = 100
	reflectType(b)

	var c *float32
	reflectType(c)

	var d myInt = 100
	reflectType(d)

	var e = person{}
	reflectType(e)
}

func TestReflect_demo1(t *testing.T) {
	var a float32 = 3.14
	reflectValue(a)

	var b int64 = 100
	reflectValue(b)

	c := person{
		name: "许志勇",
		age:  30,
	}
	reflectValue(c)
}

func TestReflect_demo2(t *testing.T) {
	var a int64 = 100
	//reflectSetValue1(a) //error : reflect.Value.SetInt using unaddressable value [recovered]
	reflectSetValue2(&a)
}

// IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic
// IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic
func TestReflect_demo3(t *testing.T) {
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())

	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
}

func TestReflect_demo4(tt *testing.T) {
	p := person{
		name: "xuzhiyong",
		age:  30,
	}
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)
	fmt.Println(fmt.Sprintf("参数个数 : %v, 方法个数 : %v", t.NumField(), t.NumMethod()))
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type() // 方法类型
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	case reflect.Struct:
		t := reflect.TypeOf(x)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			name := field.Name
			index := field.Index
			fType := field.Type
			json := field.Tag.Get("json")
			fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
				name,
				index,
				fType,
				json,
			)
		}
		// 查找指定的值
		if scoreField, ok := t.FieldByName("name"); ok {
			fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
		}
	}
}

/**
Kind 速查表
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)

**/
