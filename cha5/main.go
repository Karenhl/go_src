package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你要对自己的行为负责了")
	} else {
		fmt.Println("你还小，放过你了")
	}

}