package atomic

import (
	"errors"
	"sync/atomic"
)

type ConcurrentArray interface {
	// 用于设置指定索引上的元素值
	Set(index uint32, elem int) (err error)
	// 用于获取指定索引上的元素值
	Get(index uint32) (elem int, err error)
	// 获取长度
	Len() uint32
}

type concurrentArray struct {
	length uint32
	val    atomic.Value
}

func (array *concurrentArray) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return err
	}
	if err = array.checkValid(); err != nil {
		return err
	}
	newArray := make([]int, array.length)
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return
}

func (array *concurrentArray) Get(index uint32) (elem int, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValid(); err != nil {
		return
	}
	elem = array.val.Load().([]int)[index]
	return
}

func (array *concurrentArray) Len() uint32 {
	return array.length
}

func (array *concurrentArray) checkIndex(index uint32) error {
	if array.length <= index {
		return errors.New("invalid index")
	}
	return nil
}

func (array *concurrentArray) checkValid() error {
	if array.val.Load() == nil {
		return errors.New("Empty Array")
	}
	return nil
}

func NewConcurrentArray(length uint32) ConcurrentArray {
	array := &concurrentArray{
		length: length,
	}
	array.val.Store(make([]int, array.length))
	return array
}
