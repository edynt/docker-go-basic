package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

// func placeOrderWithoutContext(orderID string) error {
// 	log.Printf("Bắt đầu xử lý đơn hàng %s\n", orderID)

// 	const timeDelay = 3
// 	time.Sleep(timeDelay * time.Second)

// 	log.Printf(" Xử lý đơn hàng %s thành công (sau %ds)\n", orderID, timeDelay)
// 	return nil
// }

// func OrderHandler(w http.ResponseWriter, r *http.Request) {
// 	orderID := "GO-123456"
// 	err := placeOrderWithoutContext(orderID)
// 	if err != nil {
// 		http.Error(w, "Lỗi trong quá trình xử lý đơn hàng", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Đặt hàng thành công"))
// }

// // OrderHandlerSelect
// func OrderHandlerSelect(w http.ResponseWriter, r *http.Request) {
// 	orderID := "GO-123456"
// 	resultChan := make(chan error, 1)

// 	// Xử lý đơn hàng trong goroutine
// 	go func() {
// 		err := placeOrderWithoutContext(orderID)
// 		resultChan <- err
// 	}()

// 	select {
// 	case err := <-resultChan:
// 		if err != nil {
// 			log.Printf("Xử lý đơn hàng %s thất bại \n", orderID)
// 			http.Error(w, "Lỗi trong quá trình xử lý đơn hàng", http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Đặt hàng thành công"))

// 	case <-time.After(2 * time.Second):
// 		log.Printf("Xử lý đơn hàng %s quá 2s, trả lỗi về client \n", orderID)
// 		http.Error(w, "quá thời gian xử lý, vui lòng thử lại sau", http.StatusGatewayTimeout) // 504
// 	}
// }

func OrderHandlerWithContext(w http.ResponseWriter, r *http.Request) {
	orderID := "GO-123456"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := placeOrderWithContext(ctx, orderID)
	if err != nil {
		log.Printf("Xử lý đơn hàng %s thất bại %v\n", orderID, err)
		http.Error(w, "Lỗi trong quá trình xử lý đơn hàng hoặc quá thời gian", http.StatusGatewayTimeout)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Đặt hàng thành công"))
}

func placeOrderWithContext(ctx context.Context, orderID string) error {
	log.Printf("Bắt đầu xử lý đơn hàng %s\n", orderID)

	const timeDelay = 3
	select {
	case <-time.After(timeDelay * time.Second):
		log.Printf(" Xử lý đơn hàng %s thành công (sau %ds)\n", orderID, timeDelay)
		return nil
	case <-ctx.Done():
		log.Printf("Hủy xử lý đơn hàng %s: %v \n", orderID, ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// http.HandleFunc("/order", OrderHandler)
	http.HandleFunc("/order", OrderHandlerWithContext)

	log.Printf("Server listening on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
