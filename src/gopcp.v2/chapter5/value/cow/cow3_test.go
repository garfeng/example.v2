package cow

import (
	"fmt"
	"sync"
	"testing"
)

func testSetANdGet3Once(t *testing.T) {
	arr := NewConcurrentArray3(100)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			err := arr.Set(uint32(idx), idx)
			if err != nil {
				fmt.Println(err.Error())
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		item, err := arr.Get(uint32(i))
		if err != nil {
			t.Fatal(err)
		}
		if item != i {
			t.Fatalf("fail to set arr[%d] = %d", i, item)
		}
	}
}

func TestSetAndGet3(t *testing.T) {
	for i := 0; i < 100; i++ {
		testSetANdGet3Once(t)
	}
}
