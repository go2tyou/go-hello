package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 200000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		// time.Sleep(time.Millisecond * 10)
		num, ok := <- intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 { // 不是素数
				flag = false
				break
			}
		}
		if flag {
			// put primeChan 
			primeChan <- num
		}
	}
	// fmt.Println("有一个 primeNum 协程因为取不到数据，退出")
	// 这时不能关闭 primeChan
	exitChan <- true
}

func main()  {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 100000)
	exitChan := make(chan bool, 8)

	start := time.Now().Unix()
	// 写数据
	go putNum(intChan)

	// 判断符合数据
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<- exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end - start)
		// 当从 exitChan 取出 4 个结果，可放心关闭 primeChan
		close(primeChan)
	}()

	// 遍历 primeChan 取出结果
	for {
		_, ok := <- primeChan
		if !ok {
			break
		}
		// fmt.Println("res=", res)
	}

	fmt.Println("main 线程退出")
}