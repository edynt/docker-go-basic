// Bài 9: Worker Pool Xử Lý Đơn Hàng

// Yêu cầu:

// Có 10 đơn hàng.

// Tạo 3 worker goroutine đọc từ jobs channel.

// Worker xử lý đơn trong 500ms.

// Sau khi xong thì gửi "done order X" vào results.

// Main in ra kết quả.

// Gợi ý: jobs := make(chan Order, 10) + wg.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID int
}

const WORKER = 3
const ORDER = 10

func worker(id int, jobs <-chan Order, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// giả lập xử lý đơn trong 500ms
		time.Sleep(500 * time.Millisecond)
		results <- fmt.Sprintf("Worker %d: Done order %d", id, job.ID)
	}
}

func main() {
	jobs := make(chan Order, ORDER)
	results := make(chan string, ORDER)

	var wg sync.WaitGroup

	fmt.Println("call start")

	// tạo 3 worker
	for w := 1; w <= WORKER; w++ {
		wg.Add(1)
		fmt.Println("call before goroutine")
		go worker(w, jobs, results, &wg)

		fmt.Println("call after goroutine")
	}

	// gửi 10 đơn hàng vào channel jobs
	fmt.Println("call set orders into jobs")
	for o := 1; o <= ORDER; o++ {
		jobs <- Order{ID: o}
	}
	close(jobs)

	fmt.Println("call close jobs")

	// chờ tất cả worker xử lý xong
	wg.Wait()
	close(results)

	fmt.Println("call close all")

	// in ra kết quả
	for result := range results {
		fmt.Println(result)
	}
}
