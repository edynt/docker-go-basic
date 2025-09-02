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
