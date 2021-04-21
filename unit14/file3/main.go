package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	file := "F:/temp/test.txt"
	// 一次性将文件读取到位
	content, err := ioutil.ReadFile(file) // content -> []byte
	if err != nil {
		fmt.Println("read file err", err)
	}
	fmt.Printf("content is:\n%s", content)
}