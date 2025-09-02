// Bài 2: Buffered Channel

// Yêu cầu: Tạo buffered channel (size = 3). Gửi 3 số vào channel, sau đó nhận và in chúng ra.

// Gợi ý: ch := make(chan int, 3).

package main

import "fmt"

func increment(ch chan int) {
	for i := 0; i < 30; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	ch := make(chan int)
	go increment(ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println()
}
