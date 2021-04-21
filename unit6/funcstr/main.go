package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串常用系统函数
func main()  {
	// 1、统计字符串的长度，按字节 len(str)
	str1 := "hello北" // 编码统一为 utf-8 -> ascii 的字符占一个字节，汉子占 3 个字节
	fmt.Println("str1 len=", len(str1)) // str1 len= 8

	// 2、字符串的遍历，处理中文 r = []rune(str)
	str2 := "hello北京"
	str2_r := []rune(str2)
	for i := 0; i < len(str2_r); i++ {
		fmt.Printf("str2=%c\t", str2_r[i]) // str2=h  str2=e  str2=l  str2=l  str2=o  str2=北 str2=京
	}
	fmt.Println()

	// 3、字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果=", n) // 转换结果= 123
	}
	
	// 4、整数转字符串
	str4 := strconv.Itoa(1234)
	fmt.Printf("str4=%v str4=%T\n", str4, str4) // str4=1234 str4=string

	// 5、字符串转 byte 切片
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes) // bytes=[104 101 108 108 111 32 103 111]

	// 6、byte 转字符串
	str6 := string([]byte{97, 98, 99})
	fmt.Printf("str6=%v\n", str6) // str6=abc

	// 7、十进制转 2, 8, 16 进制
	str7 := strconv.FormatInt(123, 2)
	fmt.Printf("123->%v\n", str7) // 123->1111011
	str7 = strconv.FormatInt(123, 16)
	fmt.Printf("123->%v\n", str7) // 123->7b

	// 8、查找子串是否在指定字符串中
	b8 := strings.Contains("seacool", "sea")
	fmt.Printf("b8=%v\n", b8) // b8=true
	
	// 9、统计一个字符串中有几个子串
	num9 := strings.Count("sdfsfafafg", "s")
	fmt.Printf("num9=%v\n", num9) // num9=2
	
	// 10、不区分大小写的字符串比较
	b10 := strings.EqualFold("abc", "AbC")
	fmt.Printf("b10=%v\n", b10) // b10=true
	fmt.Printf("abc==AbC %v\n", "abc" == "AbC") // abc==AbC false
	
	// 11、返回子串在字符串第一次出现的 index 值，没有返回 -1
	index11 := strings.Index("hdkdoffd_939rdfjxd", "ff")
	fmt.Printf("index11=%v\n", index11) // index11=5

	// 12、返回子串在字符串最后一次出现的 index 值，没有返回 -1
	index12 := strings.LastIndex("hdkdoffd_939rdfjxd", "d")
	fmt.Printf("index12=%v\n", index12) // index12=17

	// 13、将指定的子串替换成另一个子串 最后一个参数代表换的个数，全换用 -1
	str13 := strings.Replace("go go hello", "go", "golang", -1)
	fmt.Printf("str13=%v\n", str13) // str13=golang golang hello
	
	// 14、按照指定的某个字符进行分割
	strArr14 := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr14=%v\n", strArr14) // strArr14=[hello world ok]
	for i := 0; i < len(strArr14); i++ {
		fmt.Printf("strArr14[%v]=%v\n", i, strArr14[i])
		// strArr14[0]=hello
		// strArr14[1]=world
		// strArr14[2]=ok
	}

	// 15、字符串大小写转换
	str15 := "goLang Hello"
	str15a := strings.ToLower(str15)
	fmt.Printf("str15a=%v\n", str15a) // str15a=golang hello
	str15A := strings.ToUpper(str15)
	fmt.Printf("str15A=%v\n", str15A) // str15A=GOLANG HELLO

	// 16、去除字符串左右两边空格
	str16 := " tn a lone gopher ntrn "
	str16 = strings.TrimSpace(str16)
	fmt.Printf("str16=%q\n", str16) // str16="tn a lone gopher ntrn"

	// 17、指定去除左右两边的字符
	str17 := strings.Trim("! he!llo! ", " !")
	fmt.Printf("str17=%q\n", str17) // str17="he!llo"

	// 18、以指定的字符串开头
	b18 := strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b18=%v\n", b18) // b18=true

	// 19、以指定的字符串结尾
	// strings.HasSufix("", "")
}