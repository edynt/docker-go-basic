// Bài 7: Race Condition

// Yêu cầu: Tạo biến counter = 0. 100 goroutine cùng tăng counter lên 1. In kết quả cuối cùng.

// Gợi ý: chạy 2 lần:

// - không dùng mutex (sẽ sai)

// - có mutex (sẽ đúng).

package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		wg.Add(1)

		// chi tra ra cac ket qua cuoi cung
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
