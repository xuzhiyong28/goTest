package dict

import (
	"example/project/godis/lib/utils"
	"fmt"
	"testing"
)

func TestMakeSimple(t *testing.T) {
	simpleDict := MakeSimple()
	simpleDict.m["x"] = 123
	simpleDict.m["z"] = 345
	keys := simpleDict.Keys()
	fmt.Println(keys)
}

func TestSimpleDict_Keys(t *testing.T) {
	d := MakeSimple()
	size := 10
	for i := 0; i < size; i++ {
		d.Put(utils.RandString(5), utils.RandString(5))
	}
	if len(d.Keys()) != size {
		t.Errorf("expect %d keys, actual: %d", size, len(d.Keys()))
	}
}
