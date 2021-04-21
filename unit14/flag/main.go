package main

import (
	"fmt"
	"flag"
)

func main()  {
	// 定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int
	// "u" -> -u
	// "" -> 默认值
	// "用户名，默认为空" -> 说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 转换，必须调用该方法
	flag.Parse()

	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)
}