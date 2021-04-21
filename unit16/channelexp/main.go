package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main()  {
	// var anyChan chan interface{}
	anyChan := make(chan interface{}, 10)
	anyChan <- 10
	anyChan <- "ssr"
	cat := Cat{"sm", 17}
	anyChan <- cat

	<- anyChan
	<- anyChan
	newCat := <- anyChan
	fmt.Printf("newCat=%T newCat=%v\n", newCat, newCat) // newCat=main.Cat newCat={sm 17}
	// 编译不通过 -> 使用类型断言解决
	// fmt.Printf("newCat.Name=%v\n", newCat.Name) // newCat.Name undefined (type interface {} is interface with no methods)
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name) // newCat.Name=sm
}