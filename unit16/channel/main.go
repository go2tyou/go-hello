package main

import (
	"fmt"
)

// 管道
// var 变量名 chan 数据类型
/*
 *	1、channel 是引用类型
 *	2、channel 必须初始化才能写入数据
 *	3、channel 是有类型的， intChan 只能写入整数 int
 *	4、channel 数据放满后，就不能再放入了
 *	5、如果从 channel 取出数据后，可以继续放入
 *	6、在没有使用协程的情况下，如果管道数据已经全部取出，再取就会报  deadlock!
 */
func main() {
	// 1、声明
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Println("intChan=", intChan)      // intChan= 0xc000104080
	fmt.Printf("&intChan=%p\n", &intChan) // &intChan=0xc000006028

	// 2、写
	intChan <- 10
	intChan <- 20
	num := 30
	intChan <- num
	// 当给管道写入数据时，不能超过其容量
	// intChan<- 40 // fatal error: all goroutines are asleep - deadlock!
	fmt.Printf("intChan len=%v cap=%v\n", len(intChan), cap(intChan)) // intChan len=2 cap=3

	// 取出后，可以继续放
	<-intChan // 取出扔了，可以，没毛病
	intChan <- 50

	// 3、读
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)                                        // num2= 10
	fmt.Printf("intChan len=%v cap=%v\n", len(intChan), cap(intChan)) // intChan len=2 cap=3

	// 4、在没有使用协程的情况下，如果管道数据已经全部取出，再取就会报  deadlock!
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4) // num3= 20 num4= 30
	// num5 := <-intChan // fatal error: all goroutines are asleep - deadlock!
	// fmt.Println("num5=", num5)
}
