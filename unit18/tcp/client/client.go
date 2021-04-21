package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main()  {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err...", err)
		return
	}
	fmt.Println("conn success...", conn)
	// os.Stdin -> 标准输入[终端]
	reader := bufio.NewReader(os.Stdin)
	for {
		// 从终端读取一行用户输入，并发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("conn read err...", err)
		}
		if line1 := strings.Trim(line, " \r\n"); line1 == "exit" {
			fmt.Println("client exit...")
			break
		}
		// 在将 line 发送给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn write err...", err)
		}
		// fmt.Println("client send data...", n)
	}
}