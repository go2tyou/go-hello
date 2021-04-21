package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}

func (p Phone) Start() {
	fmt.Println("phone start...")
}

func (p Phone) Stop() {
	fmt.Println("phone stop...")
}

type Camera struct {

}

func (c Camera) Start() {
	fmt.Println("Camera start...")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop...")
}

type Computer struct {

}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

// 接口
/*
 *	type 接口名 interface {
 *		method1(参数列表) 返回值列表
 *		method2(参数列表) 返回值列表
 *		...
 *	}
 *
 *	实现接口所有方法
 *
 *	func (t 自定义类型) method1(参数列表) 返回值列表 {
 *		// 方法实现
 *	}
 *
 *	func (t 自定义类型) method2(参数列表) 返回值列表 {
 *		// 方法实现
 *	}
 *
 *	// ...
 */
func main()  {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
}