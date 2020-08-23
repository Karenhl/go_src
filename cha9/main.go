package main

import (
	"fmt"
)

func main() {
	var sum int
	var count int 
	for i := 1; i <= 100; i++{
		if i % 9 == 0 {
			count += 1
			sum += i
		}
	}
	fmt.Println("count=", count, "sum=", sum)


	num2 := 6
	for num1 := 0; num1 <= 6; num1++ {
		fmt.Printf("%d + %d = %d \n", num1, num2 ,num1+num2)
		num2 = num2-1
	}
}