package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("命令行的参数", len(os.Args))
	fmt.Println("命令行的参数：")
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
}