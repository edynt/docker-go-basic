// Bài 1: Channel Hello

// Yêu cầu: Tạo 1 goroutine gửi chuỗi "Hello from goroutine" qua channel, main goroutine in ra màn hình.

// Gợi ý: make(chan string), <-ch.

package main

import "fmt"

func getHelloWorld(ch chan string) {
	ch <- "Hello world"
}

func main() {
	ch := make(chan string)
	go getHelloWorld(ch)

	fmt.Println(<-ch)
}
