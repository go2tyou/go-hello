package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	filePath := "F:/temp/abc.txt"
	file, err := os.OpenFile(filePath, os.O_TRUNC | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	defer file.Close()
	str := "xxoo,golang!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}