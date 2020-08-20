package main

import "fmt"

func main(){
	var i int = 10

	var ptr *int = &i
	fmt.Println("指针变量ptr = ", *ptr )
}
