
# Naming Conventions in Go: Short but Descriptive
> "<u>The are only two hard things in Computer Science: cache invalidation and naming things</u> - **Phil Karlton**
> (Có 2 điều khó khăn nhất trong Computer Science là *cache invalidation(Xóa và và tính toán lại khi dữ liệu thay đổi)* và *việc đặt tên*)"

> Convention sẽ tạo ra cuộc sống dễ dàng hơn nhưng những người sử dụng có thể lạm dụng nó. Điều rất quan trọng là thiết lập các convention và rule cho việc đặt tên nhưng việc áp dụng mù quáng có thể có hại nhiều hơn là lợi.

## 1. The Written Rules in Go (Các quy tắc trong Go)
### 1.1 MixedCaps
- **Nên sử dụng *MixedCaps* hoặc *mixedCaps* (đơn giản là *camelCase*) thay vì sử dụng *underscore(gạch dưới)* để đặt tên.**
- **Nếu cần truy cập từ bên ngoài package thì ký tự đầu tên của tên phải là chữ hoa như <u>M</u>ixedCaps**. Ngược lại, **nếu không cho phép sử dụng nó ở trong một package khác thì tên phải bắt đầu bằng 1 ký tự viết thường như là <u>m</u>ixedCaps .**
```go  
package awesome  
  
type Awesomeness struct {}  
// Đây là một exported method và cho phép các package khác truy cập  
func (a Awesomeness) Do() string { return a.doMagic("Awesome")  
}  
  
// Chỉ cho phép truy cập ở bên trong packge awesome  
func (a Awesomeness) doMagic(input string) string { return input  
}  
```  

- Nếu bạn gọi method doMagic từ bên ngoài -> sẽ nhận được một compile-time error(Lỗi tại thời điểm compile)

### 1.2 Interface names (Tên của các Interface)
> Interface có một method sẽ được đặt tên là: **tên method + hậu tố er** hoặc đơn giản là **chuyển các động từ thành các danh từ** như: Reader, Writer, Formatter, CloseNotifier.." - **[Go’s official documentation](https://golang.org/doc/effective_go.html)**

**=> *methodName + er = InterfaceName***

### 1. 3 Getters
- **Go không hỗ trợ tự động về setter/getter nhưng Go không cấm việc này**. Cần lưu ý một vài rule cho việc này:
> "Không có gì sai khi cung cấp getter/setter và thường sẽ thích hợp để làm như vậy nhưng không phải là phổ biến và cũng không cần thiết để đặt Get vào tên của Getter"
```go  
 owner := obj.Owner()   
 if owner != user {    
      obj.SetOwner(user)    
}  
```  
**=> Nếu một setter không thực hiện với một logic đặc biệt nào thì nên truy cập trực tiếp field.**

## 2. The Unwritten Rule in Go
> Một vài rule không phải là chính thống nhưng phổ biến trong cộng đồng.
### 2.1 Shorter variable names (Tên biến ngắn)
- **Nên sử dụng tên biến với mô tả ngắn và phải mô tả cho người đọc hiểu nội dung của nó ngay cả trước khi chạy trương trình.**
> "Chương trình phải được viết cho mọi người đọc và chỉ tình cờ cho máy móc thực thi" - Harold Abelson
- **Single-letter identifier (định danh 1 từ)** - Được đặt biệc được sử dụng trong local variable với phạm vi giới hạn. Khuyến khích Sử dụng để định danh 1 từ khi nó chỉ được giới hạn trong phạm vi của vòng lặp.
```go  
 for i := 0; i < len(pods); i++ {      //    
   }    
   ...    
   for _, p := range pods {    
     //    
   }  
```  

- **Shorthand name (Tên ngắn)** - Được khuyến khích sử dụng khi mà người khác có thể hiểu được ý nghĩa ngay từ lần đọc code đầu tiên. Phạm vi càng rộng thì tên càng cần tính mô tả.
```go
  pid // Bad (Không biết biến này nói đến podID hay PersonID hay ProductID)   
  spec // good (Nói đến Specification)    
  addr // good (Nói đến Address)  
```  

### 2.2 Unique names(Những tên riêng)
- Tên riêng như là : `API`, `HTTP`.. hoặc tên như là `ID` và `DB`. Thông thường sẽ giữ nguyên từ gốc:
  - Nên sử dụng `userID` thay vì `userId`
  - Nên sử dụng `productAPI` thay vì `productApi`

### 2.3 Line length (Độ dài)
-  Tránh việc đặt tên quá dài