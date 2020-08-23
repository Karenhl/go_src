package main
import (
	"fmt"
	"unsafe"
)

func main(){
	var a int32 = 9898
	fmt.Printf("type is %T, 占用 %d 字节", a, unsafe.Sizeof(a))
}