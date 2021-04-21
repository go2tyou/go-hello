package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	filePath1 := "F:/temp/abc.txt"
	filePath2 := "F:/temp/kakaka.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	fmt.Printf("%s", data)

	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("write file err", err)
	}
}