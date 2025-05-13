### B3: Mô phỏng 1 cái game lucky number đa luồng như sau. 
 - Lucky number là 1 số ngẫu nhiên trong khoảng [0, 9], được random bởi hệ thống mỗi 10s.
Mọi người sẽ tạo 1 cái chương trình cho phép người dùng nhập số tiền bet từ màn hình (giới hạn 1-100).
 - Mỗi lần người dùng bet số tiền x: 
> Cộng số tiền đó vào pool tiền thưởng
> In ra màn hình: System: YYYY-MM-DD HH:mm:ss: User {{i}} bet ${{x}}, the current pool is {{pool+=x}}$, waiting for {{x}}s to receive result
> Quay xổ số: Người dùng sẽ phải chờ x giây để nhận kết quả. Sau x giây, so sánh nếu x % 10 = lucky number thì thông báo người dùng trúng thưởng, trả hết số tiền trong pool hiện tại, nếu ko trúng, thông báo ko trúng.
 - In ra màn hình kết quả: 
> System: YYYY-MM-DD HH:mm:ss: User {{i}} hit the lucky number, get {{x}}$ (nếu trúng)
> System: YYYY-MM-DD HH:mm:ss: Wish user {{i}} lucky next time (nếu fail)
 - Trong lúc người dùng thứ i chờ kết quả, vẫn tiếp tục cho phép người dùng tiếp theo nhận số tiền.
 - Khi có 1 người nhập "END", dừng cho phép nhập input, chờ cho các lượt quay xổ số trước đó kết thúc và thông báo kết quả