package main

import (
	"fmt"
)


func main() {
	var num byte
	fmt.Println("请输入月份")
	fmt.Scanln(&num)

	switch num{
		case 3, 4, 5:
			fmt.Println("春天")
		case 6, 7, 8:
			fmt.Println("夏天")
		case 9, 10, 11:
			fmt.Println("秋天")
		default:
			fmt.Println("冬天")
		}
}