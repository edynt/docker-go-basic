# Các khái niệm quan trọng trong Go và cách giải quyết vấn đề

---

## 1. Goroutine
- **Là gì:** Một "luồng nhẹ" trong Go, chạy đồng thời, rẻ hơn thread.
- **Vấn đề:** Tạo quá nhiều goroutine → tốn RAM, CPU.
- **Cách giải quyết:** Dùng **Worker Pool** hoặc giới hạn goroutine qua channel.

**Ví dụ thực tế:**  
Quán cà phê có nhiều nhân viên (goroutine). Nếu thuê quá nhiều nhân viên thì lãng phí, nên chỉ cần một số lượng hợp lý.

---

## 2. Channel + `make`
- **Là gì:** Ống để các goroutine trao đổi dữ liệu.
- **Vấn đề:** Có thể bị **deadlock** nếu gửi mà không ai nhận.
- **Cách giải quyết:** Luôn đảm bảo có cả goroutine gửi và nhận. Dùng **buffered channel** nếu cần.

**Ví dụ thực tế:**  
Nhà bếp làm món xong để lên băng chuyền (channel), nhân viên phục vụ lấy ra.  
Nếu bếp đưa mà không ai lấy → tắc nghẽn.

---

## 3. Buffered Channel
- **Là gì:** Channel có bộ nhớ đệm, cho phép gửi trước nhiều dữ liệu.
- **Vấn đề:** Bộ đệm đầy mà không ai lấy → vẫn deadlock.
- **Cách giải quyết:** Thiết kế kích thước buffer hợp lý và đảm bảo luôn có goroutine nhận dữ liệu.

**Ví dụ thực tế:**  
Băng chuyền có 5 khay. Nếu không ai lấy, bếp không thể đặt thêm món nữa.

---

## 4. `select`
- **Là gì:** Cho phép chờ nhiều channel cùng lúc.
- **Vấn đề:** Nếu không có `default` hoặc timeout → có thể chờ mãi.
- **Cách giải quyết:** Dùng `time.After` hoặc `context` để timeout.

**Ví dụ thực tế:**  
Bạn chờ tin nhắn từ 2 người bạn. Nếu không ai nhắn → bạn cứ cầm điện thoại mãi.  
Giải pháp: đặt báo thức (timeout).

---

## 5. `context`
- **Là gì:** Cung cấp cơ chế hủy goroutine hoặc giới hạn thời gian.
- **Vấn đề:** Nếu không hủy → goroutine chạy mãi, gây leak.
- **Cách giải quyết:** Luôn dùng `defer cancel()`.

**Ví dụ thực tế:**  
Bạn đặt shipper giao hàng, nhưng 30 phút không thấy → bạn hủy đơn.  
Nếu không hủy → shipper cứ chạy vòng vòng.

---

## 6. `defer cancel()`
- **Là gì:** Best practice khi tạo context có timeout/cancel.
- **Vấn đề:** Nếu quên gọi → context không được giải phóng → rò rỉ tài nguyên.
- **Cách giải quyết:** Ngay sau khi tạo context thì viết `defer cancel()`.

**Ví dụ thực tế:**  
Bật máy bơm (context). Nếu không tắt (cancel) → tốn điện, hỏng máy.

---

## 7. Mutex
- **Là gì:** Dùng để bảo vệ dữ liệu chung khi nhiều goroutine cùng đọc/ghi.
- **Vấn đề:** Quên unlock → deadlock. Lạm dụng → giảm hiệu năng.
- **Cách giải quyết:** Luôn dùng `defer mu.Unlock()`. Dùng `atomic` nếu chỉ cần thao tác đơn giản.

**Ví dụ thực tế:**  
Có 1 cuốn sổ chung. Ai muốn ghi phải mượn chìa khóa (mutex). Nếu giữ mãi chìa khóa → người khác không ghi được.

---

## 8. `sync.Once`
- **Là gì:** Đảm bảo một đoạn code chỉ chạy một lần duy nhất.
- **Vấn đề:** Nếu quên dùng → có thể chạy nhiều lần không cần thiết.
- **Cách giải quyết:** Dùng `once.Do(fn)` cho code init hoặc setup.

**Ví dụ thực tế:**  
Chỉ cần bật server 1 lần. Nếu nhiều người cùng bấm nút → `sync.Once` sẽ chặn.

---

## 9. WaitGroup
- **Là gì:** Dùng để chờ nhiều goroutine kết thúc.
- **Vấn đề:** Nếu quên gọi `wg.Done()` → chương trình chờ mãi.
- **Cách giải quyết:** Luôn gọi `defer wg.Done()` trong goroutine.

**Ví dụ thực tế:**  
Bạn tổ chức tiệc, chờ bạn bè dọn bàn ghế xong mới ăn. Nếu 1 người không 
báo xong → bạn cứ chờ.

---

## 10. `recover()` và `panic()`
- **Là gì:** `panic` làm chương trình crash, `recover` giúp bắt lỗi và tiếp tục chạy.
- **Vấn đề:** Không recover → chương trình sập.
- **Cách giải quyết:** Dùng `recover()` trong `defer` để xử lý an toàn.

**Ví dụ thực tế:**  
Xe đang chạy thì nổ lốp (panic). Nếu có bánh dự phòng (recover) → thay tiếp tục đi. Nếu không → đứng im.

---

## 11. Interface
- **Là gì:** Định nghĩa hành vi (method) mà struct cần có.
- **Vấn đề:** Nếu thiết kế interface quá to → khó implement.
- **Cách giải quyết:** Dùng interface nhỏ (interface segregation).

**Ví dụ thực tế:**  
Interface “biết chạy” → chỉ cần method `Run()`. Xe máy, ô tô, xe bus… đều dùng được.

---

## 12. Pointer
- **Là gì:** Làm việc với địa chỉ bộ nhớ, thay đổi trực tiếp dữ liệu gốc.
- **Vấn đề:** Nếu không cẩn thận → thay đổi dữ liệu ngoài ý muốn.
- **Cách giải quyết:** Dùng pointer khi cần hiệu năng, tránh copy dữ liệu lớn.

**Ví dụ thực tế:**  
Đưa bản photo chìa khóa (copy) → sửa không ảnh hưởng chìa thật.  
Đưa chìa thật (pointer) → mở cửa trực tiếp.

---

## 13. `sync/atomic`
- **Là gì:** Cho phép thao tác an toàn trên biến số nguyên mà không cần Mutex.
- **Vấn đề:** Chỉ dùng cho số nguyên đơn giản, không thay được Mutex trong mọi tình huống.
- **Cách giải quyết:** Dùng `atomic.AddInt32`, `atomic.LoadInt32` khi chỉ cần đếm hoặc cộng dồn.

**Ví dụ thực tế:**  
Máy phát số thứ tự trong bệnh viện. Ai bấm nút → số tăng 1 cách an toàn.

---

## 14. Worker Pool
- **Là gì:** Mẫu thiết kế dùng một số lượng goroutine cố định xử lý nhiều công việc.
- **Vấn đề:** Nếu không giới hạn → có thể tạo quá nhiều goroutine, gây nghẽn.
- **Cách giải quyết:** Dùng worker pool để phân phối việc hợp lý.

**Ví dụ thực tế:**  
Có 100 đơn hàng. Thay vì thuê 100 shipper → chỉ cần 5 shipper, họ lấy đơn lần lượt từ kênh.

---

## 15. Goroutine Leak
- **Là gì:** Goroutine được tạo ra nhưng không bao giờ kết thúc.
- **Vấn đề:** Rò rỉ tài nguyên, tốn RAM.
- **Cách giải quyết:** Dùng `context` để hủy, đóng channel khi không dùng nữa, hoặc dùng `select` với timeout.

**Ví dụ thực tế:**  
Bạn thuê shipper nhưng không bao giờ hủy đơn. Shipper cứ chạy mãi ngoài đường → tốn xăng.

---
