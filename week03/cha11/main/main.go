package main

import (
	"fmt"
	utils "go_src/week03/cha11/utils"
)

func main() {
	var num int
	fmt.Println("input your index:")
	fmt.Scanln(&num)
	res := utils.Test(num)
	fmt.Println("num=",res)
}