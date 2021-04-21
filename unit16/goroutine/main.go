package main

import (
	"fmt"
	"strconv"
	"time"
)

func test()  {
	for i := 0; i < 10; i++ {
		fmt.Println("hello,world" + strconv.Itoa(i + 1))
		time.Sleep(time.Second)
	}	
}

/*
 *	1、主线程是一个物理线程，直接作用在 CPU 上，是重量级的，非常耗费 CPU 资源
 *	2、协程从主线程开启的，是轻量级的线程，是逻辑态，对资源消耗相对较小
 *	3、golang 的协程机制是重要的特点，可以轻松开启上万个协程，突出 go 的并发优势
 */
func main()  {
	go test() // 开启了一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("hello,golang" + strconv.Itoa(i + 1))
		time.Sleep(time.Second)
	}
}

// 以下说明 main 主线程和 test 协程同时执行
// hello,golang1
// hello,world1
// hello,world2
// hello,golang2
// hello,world3
// hello,golang3
// hello,golang4
// hello,world4
// hello,world5
// hello,golang5
// hello,world6
// hello,golang6
// hello,world7
// hello,golang7
// hello,golang8
// hello,world8
// hello,world9
// hello,golang9
// hello,golang10
// hello,world10