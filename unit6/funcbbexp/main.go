package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func (string) string {
	return func (name string) string {
		// 判断某个字符串有没有指定的后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix	
		}
		return name
	}
}

func main()  {
	f := makeSuffix(".jpg")
	fmt.Println("处理后=", f("winter"))
	fmt.Println("处理后=", f("spring.jpg"))
}