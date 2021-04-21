package main

import (
	"fmt"
	"errors"
)

func test()  {
	// 使用 defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover() 内置函数，可以捕获异常
		if err != nil {
			fmt.Println("err=", err)
			/// 发送通知
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取错误。。。")
	}
}

func test2()  {
	err := readConf("configure.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test2...")
}

// 1、默认情况下，当错误(panic)，程序就会退出
// 2、希望当发生错误后，可以捕获到错误并进行处理，保证程序可以继续执行，错误通知
// 3、defer panic recover
// 		go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后处理
// 4、自定义错误
// 		errors.New("xxx") + panic() 
func main()  {
	// test()
	// fmt.Println("main...")

	test2()
	fmt.Println("main...")
}