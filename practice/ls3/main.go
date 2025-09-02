// Bài 3: Timeout

// Yêu cầu: Viết 1 hàm giả lập API chậm (sleep 2s). Main goroutine chỉ chờ 1s → nếu không có kết quả thì in "timeout".

// Gợi ý: select { case <-time.After(...) }.

// Cách giải:
// Tạo một channel ch để nhận kết quả từ goroutine giả lập API.

// Trong goroutine: time.Sleep(2 * time.Second) rồi gửi "API done" vào channel.

// Ở main, dùng select:

// Nếu nhận được dữ liệu từ ch → in kết quả.

// Nếu time.After(1 * time.Second) hết hạn → in "timeout".

package main

import (
	"fmt"
	"time"
)

func slowApi(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "API done"
}

func main() {
	ch := make(chan string)

	go slowApi(ch)

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
