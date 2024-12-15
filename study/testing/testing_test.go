package testing

import (
	"fmt"
	"testing"
)

func sum(x, y int) int {
	return x + y
}

// 子测试
func TestChildRun(t *testing.T) {
	testCalss := []struct {
		gmt  string
		loc  string
		want string
	}{
		{"12:31", "Europe/Zuri", "13:31"},     // incorrect location name
		{"12:31", "America/New_York", "7:31"}, // should be 07:31
		{"08:08", "Australia/Sydney", "18:08"},
	}
	for _, tc := range testCalss {
		t.Run(fmt.Sprintf("%s_%s", tc.gmt, tc.loc), func(t *testing.T) {
			t.Logf("%s , %s , %s", tc.gmt, tc.loc, tc.want)
		})
	}
}

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

func TestSlice(t *testing.T) {
	p1 := NewPerson("小王子", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}
