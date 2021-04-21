package main

import (
	"fmt"
	"os"
)

func main()  {
	// file 叫法
	// file 对象 | file 指针 | file 文件句柄
	file, err := os.Open("F:/temp/test.txt")
	if err != nil {
		fmt.Println("open file err", err)
	}
	fmt.Printf("file=%v", file) // file=&{0xc000080780} -> file 就是一个指针 *File
	err1 := file.Close()
	if err1 != nil {
		fmt.Println("close file err1", err1)
	}
}