// Bài 8: Inventory Stock

// Yêu cầu: Tạo struct Inventory{Stock int, mu sync.Mutex}. 5 goroutine cùng mua hàng (trừ stock). Đảm bảo stock không âm.

// Gợi ý: Lock()/Unlock().
package main

import (
	"fmt"
	"sync"
)

type Inventory struct {
	Stock int
	mu    sync.Mutex
}

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
	inv := &Inventory{Stock: 10}

	// Đợi đến khi xong
	var wg sync.WaitGroup

	// Danh sách người mua
	buyers := []string{"A", "B", "C", "D", "E"}

	for _, name := range buyers {
		wg.Add(1)

		// sử dụng goroutine là để mua hàng cùng lúc, ai mua được thì mua chứ không phải mua tuần tự
		go func(buyer string) {
			defer wg.Done()
			success := inv.Buy(3, buyer)
			if success {
				fmt.Printf("✔ %s mới mua hàng\n", buyer)
			} else {
				fmt.Printf("✘ %s không thể mua hàng\n", buyer)
			}
		}(name)
	}

	wg.Wait()
	fmt.Printf("Số hàng còn lại: %d\n", inv.Stock)

}
