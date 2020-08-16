package main

import (
    "hash/crc32"
	"io"
	"fmt"
)

func CalCrc32(checkStr string) uint32 {
    // checkStr: 待计算的字符串
    iEe := crc32.NewIEEE()
    io.WriteString(iEe, checkStr)
    return iEe.Sum32()
}


func main() {
	str1 := CalCrc32("4c0ee35bb7547c80d0488c437e7899d8")
	fmt.Println(str1)
}


