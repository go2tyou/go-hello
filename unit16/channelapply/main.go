package main

import (
	"fmt"
	"time"
)

func WriteData(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write data=", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(intChan)
}

func ReadData(intChan chan int, exitChar chan bool)  {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Println("read data=", v)
		time.Sleep(time.Millisecond * 100)
	}
	// read finshed
	exitChar <- true
	close(exitChar)
}

// 1、如果编译器（运行），发现一个通道只有写，而没有读，则该管道会阻塞
// 2、写管道与读管道的频率不一致，无所谓
func main()  {
	// 创建两个管道
	intChar := make(chan int, 10)
	exitChar := make(chan bool, 1)
	go WriteData(intChar)
	go ReadData(intChar, exitChar)

	for {
		_, ok := <- exitChar
		if !ok {
			break
		}
	}
}