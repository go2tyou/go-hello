package main

import (
	"fmt"
	"runtime"
)

func main()  {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu num is", cpuNum) // cpu num is 8

	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}