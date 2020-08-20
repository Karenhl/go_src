package main
import (
	"fmt"
)

func main() {
	// 声明2个int32变量并赋值，判断2个之和，如果大于等于50输出helloworld
	var a int32 = 10
	var b int32 = 50
	if a + b >= 50 {
		fmt.Println("hello world")
	}

	// 声明2个float64变量，判断第一个数大于10.0，且第二个数小于20.0
	var n1 float64 = 11.0
	var n2 float64 = 17.0
	if n1 > 10.0 && n2 < 20.0 {
		fmt.Println("和等于", n1+n2)
	}

	// 定义2个int32变量，判断二者之和是否能被3和5整除，打印提示信息
	var n3 int32 = 1
	var n4 int32 = 9
	if (n3 + n4) % 3 == 0 && (n3 + n4) % 5 == 0{
		fmt.Println("可以")
	} else {
		fmt.Println("不可以")
	}

	// 判断是否是闰年（闰年，符合条件之一：（1）年份能被4整除，但不能被100整除 （2）能被400整除）
	var year int
	fmt.Println("请输入年份")
	fmt.Scanln(&year)

	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		fmt.Println("是闰年")
	} else {
		fmt.Println("不是闰年")
	}
}