package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	// 循环接收客户端发送的数据
	defer conn.Close()
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1、等待客户端通过 conn 发送信息
		// 2、如果客户端没有 write[发送]，那么协程就阻塞在这里
		// fmt.Println("server waiting client ", conn.RemoteAddr().String(), " send msg...")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server read err...", err)
			return
		}
		// 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n])) // buf[:n] -> 实际读到的内容
	}
}

func main()  {
	fmt.Println("server start listen...")
	// tcp -> 使用网络协议是 tcp
	// 0.0.0.0:8888 -> 本地监听 8888 端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err...", err)
		return
	}
	// fmt.Printf("listen suc=%v\n", listen) // listen suc=&{0xc0000cea00 {<nil> 0}}
	defer listen.Close() // 延时关闭
	// 循环等待客户端来连接
	for {
		fmt.Println("waiting client conn...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err...", err)
		} else {
			fmt.Printf("accept suc conn=%v client ip=%v\n", conn, conn.RemoteAddr().String())
		}
		// 准备启动一个协程，为客户端服务
		go process(conn)
	}
}