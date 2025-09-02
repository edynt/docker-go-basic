// Bài 5: Cancel

// Yêu cầu: Tạo 1 goroutine chạy vòng lặp in "working..." mỗi 200ms. Main goroutine cancel sau 1s → worker dừng.

// Gợi ý: ctx, cancel := context.WithCancel(...).

package main

import (
	"context"
	"log"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Worker is done")
			return
		default:
			log.Println("Working...")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()
}
