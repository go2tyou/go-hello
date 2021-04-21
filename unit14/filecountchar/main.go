package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

func main()  {
	filePath := "F:/temp/xxx.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("open flie err", err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
    for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// str1 := []rune(str) // 中文
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Println(count)	
}