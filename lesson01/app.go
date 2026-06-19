package main
import "fmt"
func main() {
	// fmt.Println("Hello World")
	// randomUser();
	var fullName string = "Nguyễn Phúc Thịnh";
	fmt.Println(fullName);
	var age int;
	// Nếu không khai báo giá trị thì mặc định bằng 0
	age=20;
	fmt.Println(age);
	// := khai báo giá trị ko cần khai báo kiểu
	//Lưu ý: chỉ dùng được trong 1 function
	phone := "0123456789";
	fmt.Println(phone);
	toan, van, anh := 10, 20, 30;
	fmt.Println(toan, van, anh);
}