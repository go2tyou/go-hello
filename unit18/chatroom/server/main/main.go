package main

import (
	"fmt"
	"net"
	"time"
	"gocode/helloworld/unit18/chatroom/server/model"
)

// 处理和客户端的通讯
func process(conn net.Conn)  {
	// 延时关闭
	defer conn.Close()
	
	// 读取客户端发送的信息
	fmt.Println("读取客户端发送的数据......")
	// 创建一个 Processor
	processor := &Processor {
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误......", err)
		return
	}
}

// 对 UserDao 初始化
func initUserDao()  {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	initPool("0.0.0.0:6379", 16, 0, 300 * time.Second)
	initUserDao()
}

func main()  {
	fmt.Println("服务器在 8889 端口监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err......", err)
		return
	}
	// 监听成功，等待客户端来连接
	for {
		fmt.Println("等待客户端来连接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err......", err)
		}
		// 连接成功，启动一个协程和客户端保持通讯
		go process(conn)
	}
}