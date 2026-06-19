package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	var hoten string;
	fmt.Print("Nhap ho va ten: ");
	//fmt.Scanln(&hoten);
	scanner := bufio.NewScanner(os.Stdin);
	if (scanner.Scan()) {
	hoten = scanner.Text();
	}
	fmt.Println("Xin chao: ", hoten);
}