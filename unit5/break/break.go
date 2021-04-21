package main

import (
	"fmt"
	"math/rand"
	"time"
)

// break 用于终止某个语句块的执行，用于中断 for 循环或退出 switch 语句
// 在多层嵌套的语句块中，可以通过 标签(label) 指明要终止的是那一层语句块
func main()  {

	// time.Now().Unix() ：返回 1970.01.01 00:00:00 到现在的秒数
	// rand.Seed(time.Now().Unix())
	// n := rand.Intn(100) + 1 // [0, 100)
	// fmt.Println("n=", n)

	var count int = 0
	for {
		// rand.Seed(time.Now().Unix())
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(100) + 1
		count++
		if x == 99 {
			break
		}
	}
	fmt.Printf("生成 99 一共使用了 %v 次\n", count)

	label2:
	for i := 0; i < 4; i++ {
		// label1: // 设置一个标签，名字可自定义
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break // 默认跳出最近的 for 循环
				// break label1 // 可以指定标签，跳出标签对应的 for 循环
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}