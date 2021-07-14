package dict

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestConcurrentPut(t *testing.T) {
	d := MakeConcurrent(0)
	fmt.Println(d)
	count := 100
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0 ; i < count ; i++ {
		go func(i int) {
			key := "k" + strconv.Itoa(i)
			ret := d.Put(key, i)
			if ret != 1 {
				t.Error("put test failed: expected result 1, actual: " + strconv.Itoa(ret) + ", key: " + key)
			}
			val, ok := d.Get(key)
			if ok {
				intVal, _ := val.(int)
				if intVal != i {
					t.Error("put test failed: expected " + strconv.Itoa(i) + ", actual: " + strconv.Itoa(intVal) + ", key: " + key)
				}
			} else {
				_, ok := d.Get(key)
				t.Error("put test failed: expected true, actual: false, key: " + key + ", retry: " + strconv.FormatBool(ok))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
