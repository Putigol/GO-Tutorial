package main

import "fmt"

const deliciousDish="Bánh mì"
func main() {
	//Hằng số không thể ghi đè trong cùng 1 phạm vi
	const deliciousDish="Cơm tấm"
	fmt.Println("Món ngon nhất là: ",deliciousDish)

}