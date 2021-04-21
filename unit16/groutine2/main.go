package main

import (
	"fmt"
	// "time"
	"sync"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i < n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

// go build -race .\main.go
func main()  {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
    // fatal error: concurrent map writes
	// 休眠 10s ？
	// time.Sleep(time.Second * 10)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%v\n", i, v)
	}
	lock.Unlock()
}