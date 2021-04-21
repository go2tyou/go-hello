package main

import (
	"fmt"
	"time"
)

func sayHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

func test()  {
	// defer + recover
	defer func() {
		// 捕获 panic
		if err := recover(); err != nil {
			fmt.Println("test() err", err) // test() err assignment to entry in nil map
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" // err
	// panic: assignment to entry in nil map
}

/*
	1、goroutine 中使用 recover，解决协程中出现 panic，导致程序崩溃
*/
func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("hello,go")
		time.Sleep(time.Second)
	}
}
