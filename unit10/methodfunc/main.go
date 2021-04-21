package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func test01(p Person)  {
	fmt.Println("test01()...", p.Name)
}

func test02(p *Person)  {
	fmt.Println("test02()...", p.Name)
}

func (p Person) test03()  {
	p.Name = "sr"
	fmt.Println("test03()...", p.Name)
}

func (p *Person) test04()  {
	p.Name = "r"
	fmt.Println("test04()...", p.Name)
}

// 1、对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
// 2、对于方法，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来也同样可以
// 3、关键是：方法和哪个类型绑定
func main()  {
	var p Person
	p.Name = "ssr"
	test01(p) // test01()... ssr
	test02(&p) // test02()... ssr

	p.test03()
	fmt.Println("main()...", p.Name) // test03()... sr main()...... ssr
	(&p).test03() // 形式上传入了地址，本质仍然是值拷贝
	fmt.Println("main()......", p.Name) // test03()... sr main()...... ssr

	(&p).test04()
	fmt.Println("main()......", p.Name) // test04()... r main()...... r
	p.test04() // 形式上传入值类型，本质是地址拷贝
	fmt.Println("main()......", p.Name) // test04()... r main()...... r
}