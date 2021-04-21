package main

import (
	"fmt"
)

func main()  {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan)
	// 管道关闭后不能写，只能读
	// 写
	// intChan <- 300 // panic: send on closed channel
	// 读
	n1 := <- intChan
	fmt.Println("n1=", n1) // n1= 100

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// len(intChan2) 不断变化，不能使用此方式 
	// for i := 0; i < len(intChan2); i++ {
	// }

	// 1、这时如果 channel 没有关闭，则会出现 deadlock!
	// fatal error: all goroutines are asleep - deadlock!	

	// 2、这时如果 channel 关闭，则正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}