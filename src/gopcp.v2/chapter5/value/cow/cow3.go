package cow

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type int64Array struct {
	length uint32
	val    atomic.Value
}

func NewConcurrentArray3(length uint32) ConcurrentArray {
	arr := int64Array{}
	arr.length = length
	arr.val.Store(make([]int64, arr.length))
	return &arr
}

func (array *int64Array) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}

	elemInt64 := int64(elem)
	atomic.StoreInt64(&(array.val.Load().([]int64)[index]), elemInt64)
	return
}

func (array *int64Array) Get(index uint32) (elem int, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}

	elem = int(array.val.Load().([]int64)[index])
	return
}

func (array *int64Array) Len() uint32 {
	return array.length
}

// checkIndex 用于检查索引的有效性。
func (array *int64Array) checkIndex(index uint32) error {
	if index >= array.length {
		return fmt.Errorf("Index out of range [0, %d)!", array.length)
	}
	return nil
}

// checkValue 用于检查原子值中是否已存有值。
func (array *int64Array) checkValue() error {
	v := array.val.Load()
	if v == nil {
		return errors.New("Invalid int array!")
	}
	return nil
}
