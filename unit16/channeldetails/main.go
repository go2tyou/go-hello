package main

import (
	"fmt"
)

/*
	1、channel 可以定义为只读或只写（默认双向）
*/
func main() {
	// var chan1 chan int // 可读可写

	// 只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	fmt.Println("chan2", chan2)
	// num := <- chan2 // invalid operation: <-chan2 (receive from send-only type chan<- int)

	// 只读
	var chan3 <-chan int
	num2 := <-chan3
	// chan3 <- 30 // invalid operation: chan3 <- 30 (send to receive-only type <-chan int)
	fmt.Println("num2", num2)
}
