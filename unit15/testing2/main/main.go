package main

// 细节
/*
 *	1、测试用例文件名必须以 _test.go 结尾，
 *	2、测试用例函数必须 TestXxx
 *	3、测试用例函数形参类型必须是 t *testing.T
 *	4、一个测试用例函数中，可以有多个测试用例函数
 *	5、运行测试用例指令 go test -v(运行正确错误都输出日志)
 *	6、当出现错误时，使用 t.Fatalf()
 *	7、输出日志 t.Logf()
 *	8、PASS 测试用例运行成功 FAIL 反之
 *	9、测试单个文件 go test -v cal_test.go cal.go
 *	10、测试单个方法 go test -v -test.run TestAddUpper
 */
func main()  {
	
}