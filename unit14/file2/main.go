package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main()  {
	file, err := os.Open("F:/temp/test.txt")
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer file.Close() // 及时关闭，否则会有内存泄漏
	// 创建一个 *Reader 是带缓冲的 const (defaultBufSize = 4096)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Print("file read end...")
}