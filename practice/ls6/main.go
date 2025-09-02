// Bài 6: Timeout API

// Yêu cầu: Call một API công khai (vd: https://httpbin.org/delay/3). Dùng context.WithTimeout set 2s. Nếu API chậm hơn → in "request timeout".

// Gợi ý: http.NewRequestWithContext.

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	//  tạo context với timeout 2s
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// tạo request với context.
	// call api delay with greater than 2s
	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		fmt.Println("Erro creating request:", err)
		return
	}

	//  tạo client và gửi request
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// đọc response nếu thành công
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}
	fmt.Println("Response:", string(bodyBytes))
}
