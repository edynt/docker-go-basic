// Bài 4: Multi-channel

// Yêu cầu: Tạo 2 goroutine, một gửi "ping" mỗi 500ms, một gửi "pong" mỗi 700ms. Main goroutine dùng select để in ra "ping" hoặc "pong" khi nhận được.

// Gợi ý: vòng for { select { ... } }.

package main

import (
	"fmt"
	"time"
)

func getPing(ch chan string) {
	for {
		ch <- "ping 500"
		time.Sleep(500 * time.Millisecond)
	}
}

func getPong(ch chan string) {
	for {
		ch <- "pong 700"
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	ch := make(chan string)

	go getPing(ch)
	go getPong(ch)

	for {
		select {
		case result := <-ch:
			fmt.Println(result)
		default:
			time.Sleep(10 * time.Millisecond)
		}

	}
}
