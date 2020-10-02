package cow

import (
	"fmt"
	"sync"
	"testing"
)

func testSetAndGet3Once(n int, t *testing.T) {
	arr := NewConcurrentArray3(uint32(n))
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
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

	for i := 0; i < n; i++ {
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
	// 1000 次测试
	for i := 0; i < 1000; i++ {
		// 10000 并发
		testSetAndGet3Once(10000, t)
	}
}
