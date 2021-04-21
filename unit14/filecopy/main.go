package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func copyFile(dstFileName string, srcFileName string) (writtn int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err", err)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)
	return io.Copy(writer, reader)
}

func main()  {
	filePath1 := "C:/Users/minso/Desktop/ide-eval-resetter-2.1.6.zip"
	filePath2 := "F:/temp/xyz.zip"
	_, err := copyFile(filePath2, filePath1)
	if err == nil {
		fmt.Println("copy fished")
	} else {
		fmt.Println("copy err", err)
	}
}