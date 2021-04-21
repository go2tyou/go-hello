package main

import (
	"fmt"
)

/*
	1、使用 select 可以解决从管道取数据的阻塞问题
*/
func main() {
	intChan := make(chan int, 10) 
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	strChan := make(chan string, 5) 
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统方法在遍历管道时，如果不关闭会阻塞而导致 deadlock!
	// 有时候，不方便确定什么时候关闭管道
	// 使用 select 解决
	// label:
	for {
		select {
		// 注意：这里，如果 inChan 一直没有关闭，不会一直阻塞而 deadlock!
		// 会自动到下一个 case 匹配	
		case v := <-intChan:
			fmt.Println("from intChan read data", v)
		case v := <-strChan:
			fmt.Println("from strChan read data", v)
		default:
			fmt.Println("read data nonono~")
			// 加入业务逻辑
			return
			// break label
		}
	}
}
