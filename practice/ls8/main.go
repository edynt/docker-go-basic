// Bài 8: Inventory Stock

// Yêu cầu: Tạo struct Inventory{Stock int, mu sync.Mutex}. 5 goroutine cùng mua hàng (trừ stock). Đảm bảo stock không âm.

// Gợi ý: Lock()/Unlock().
package main

import (
	"fmt"
	"sync"
	"time"
)

// Inventory đại diện cho kho hàng
type Inventory struct {
	Stock int
	mu    sync.Mutex
}

// Mua hàng (giảm stock an toàn với Mutex)
// Trả về true nếu mua được, false nếu hết hàng
func (inv *Inventory) Buy(item int, buyer string) bool {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if inv.Stock >= item {
		inv.Stock -= item
		fmt.Printf("%s mua thành công %d sản phẩm\n", buyer, item)
		return true
	}

	fmt.Printf("%s không mua được (hết hàng)\n", buyer)
	return false
}

func main() {
	inv := &Inventory{Stock: 10} // kho ban đầu có 10 sản phẩm
	var wg sync.WaitGroup

	// Danh sách người mua
	buyers := []string{"A", "B", "C", "D", "E"}

	for _, name := range buyers {
		wg.Add(1)
		go func(buyer string) {
			defer wg.Done()
			success := inv.Buy(3, buyer)
			time.Sleep(200 * time.Millisecond)
			if success {
				fmt.Printf("✔ %s đã mua hàng thành công\n", buyer)
			} else {
				fmt.Printf("✘ %s không thể mua hàng\n", buyer)
			}
		}(name)
	}

	wg.Wait()
	fmt.Printf("Số hàng còn lại: %d\n", inv.Stock)
}
