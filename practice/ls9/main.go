// Bài 9: Worker Pool Xử Lý Đơn Hàng

// Yêu cầu:

// Có 10 đơn hàng.

// Tạo 3 worker goroutine đọc từ jobs channel.

// Worker xử lý đơn trong 500ms.

// Sau khi xong thì gửi "done order X" vào results.

// Main in ra kết quả.

// Gợi ý: jobs := make(chan Order, 10) + wg.

package ls9
