// Bài 10: Checkout Service (E-Commerce Simulation)

// Yêu cầu:

// Order gồm ID, Amount.

// Khi checkout:

// Trừ tồn kho (có mutex).

// Gọi API thanh toán (fake bằng time.Sleep). Có timeout.

// Nếu thành công → gửi job "send email" vào channel.

// Nếu timeout hoặc cancel → rollback stock.

// Gợi ý:

// context.WithTimeout cho thanh toán.

// mutex cho stock.

// channel cho job email.

// select để xử lý timeout.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID     string
	Amount int
}

var stock = 100
var mu sync.Mutex

// Giả lập thanh toán có timeout
func processPayment(ctx context.Context, order Order) error {
	ch := make(chan bool, 1)

	fmt.Println("Start: Before goroutine processing payment for", order.ID)

	go func() {
		fmt.Println("Goroutine Processing payment for", order.ID)
		// giả lập call API thanh toán mất 2s
		time.Sleep(2 * time.Second)
		ch <- true // báo thành công
	}()

	fmt.Println("Start: After goroutine processing payment for", order.ID)

	select {
	case <-ctx.Done():
		return fmt.Errorf("payment timeout or canceled")
	case <-ch:
		fmt.Println("Payment success for", order.ID)
		return nil // thành công
	}
}

func checkout(order Order, emailJobs chan<- string) {
	fmt.Println("Start function checkout")
	// 1. Lock tồn kho
	mu.Lock()
	if stock < order.Amount {
		mu.Unlock()
		fmt.Println("❌ Not enough stock for", order.ID)
		return
	}
	stock -= order.Amount
	mu.Unlock()

	// 2. Context timeout 3s cho thanh toán
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println("Checkout :: call before call process payment")
	err := processPayment(ctx, order)
	fmt.Println("Checkout :: processPayment done")
	if err != nil {
		// rollback nếu fail
		mu.Lock()
		stock += order.Amount
		mu.Unlock()
		fmt.Println("⚠️ Payment failed for", order.ID, "-> rollback")
		return
	}

	// 3. Nếu thành công -> gửi job email
	fmt.Println("Call set message to emailJobs")
	emailJobs <- fmt.Sprintf("📧 Order %s confirmed!", order.ID)
	fmt.Println("✅ Checkout success for", order.ID)
}

func emailWorker(jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Sending email:", job)
		time.Sleep(1 * time.Second) // giả lập gửi email
	}
}

func main() {
	emailJobs := make(chan string, 10)
	var wg sync.WaitGroup

	// tạo 1 worker xử lý email
	wg.Add(1)
	go emailWorker(emailJobs, &wg)

	// tạo một số order để test
	orders := []Order{
		{ID: "ORDER_0001", Amount: 10},
		{ID: "ORDER_0002", Amount: 50},
		{ID: "ORDER_0003", Amount: 60}, // sẽ fail vì stock không đủ
	}

	for _, order := range orders {
		checkout(order, emailJobs)
	}

	fmt.Println("Call close emailJobs")
	close(emailJobs) // đóng channel sau khi không còn order
	wg.Wait()

	fmt.Println("📦 Final stock:", stock)
}
