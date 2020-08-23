package main

import (
	"fmt"
)



func main(){


/*

   *
  * *
 *****

*/
	var layer int = 19
	
	// 打印的层数
	for i := 1; i <= layer; i++ {
		// 每层的内容
		for k := 1; k <= layer - i; k++{
			fmt.Print(" ")
		}
		for j := 1; j <= i * 2 - 1; j++{
			if j == 1 || j == i * 2 - 1||i == layer{
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}

/*
	打印九九乘法表
*/
	layer2 := 9
	for i := 1 ; i <= layer2; i++{
		for j := 1; j <= i ; j++{
			fmt.Printf("%v * %v = %v \t", i, j, i*j)
		} 
		fmt.Println()
	}

}