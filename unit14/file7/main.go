package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main()  {
	filePath := "F:/temp/abc.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", str)
	}

	str := "aaaaaa!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}