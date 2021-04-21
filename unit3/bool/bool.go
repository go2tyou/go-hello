package main

import (
	"fmt"
	"unsafe"
)

/**
 * 1、只允许 true 和 false
 * 2、占一个字节
 */

// go 中布尔类型的使用
func main() {
	var b = false
	fmt.Println("b=", b)
	fmt.Printf("b 占用空间=%d", unsafe.Sizeof(b))
}
