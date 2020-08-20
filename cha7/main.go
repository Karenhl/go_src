package main
import (
	"fmt"
)

func main() {
	var score float64
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	if score == 100.0 {
		fmt.Println("获得BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("iphone8")
	} else if score > 60 && score <= 80 {
		fmt.Println("ipad")
	} else {
		fmt.Println("什么都没有")
	}
}