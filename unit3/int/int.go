package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)

/**
 * 1、有符号和无符号
 * 2、整型默认 int
 * 3、整型变量在使用时，遵循保小不保大的原则（在保证程序正常运行下，尽量使用占用空间小的数据类型）
 * 		var age byte = 90
 * 4、1 byte = 8 bit
 */

// go 中整数类型的使用
func main() {

	var i int = 1
	fmt.Println("i=", i)

	// 测试 int8 的范围 -128 ~ 127
	var j int8 = -128
	// var j int8 = 129
	// var j int8 = -129 // constant -129 overflows int8
	fmt.Println("j=", j)

	// 测试 uint8 的范围 0 ~ 255
	var k uint8 = 0
	// var k uint8 = -1 // constant -1 overflows uint8
	// var k uint8 = 256
	fmt.Println("k=", k)

	// int uint rune byte 的使用
	var a int = 8888
	fmt.Println("a=", a)
	
	var b uint = 2
	fmt.Println("b=", b)
	
	var c byte = 255
	fmt.Println("c=", c)

	var n1 = 100 // n1 是什么类型
	fmt.Printf("n1 的类型是 %T \n", n1) // n1 的类型是 int

	// 某个变量占用字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T，n2 占用的字节数 %d", n2, unsafe.Sizeof(n2)) // n2 的类型 int64，n2 占用的字节数 8
}
