package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("phone start...")
}

func (p Phone) Stop() {
	fmt.Println("phone stop...")
}

func (p Phone) Call() {
	fmt.Println("phone call...")
}

type Camera struct {
	Name string
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
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main()  {
	// 定义一个 Usb 接口数组，可以存放 Phone 和 Camera 结构体
	// 这里体现出多态数组
	var usbArr [3]Usb
	fmt.Println(usbArr) // [<nil> <nil> <nil>]
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"mi"}
	usbArr[2] = Camera{"v3d"}
	fmt.Println(usbArr) // [{vivo} {mi} {v3d}]

	var computer Computer
	for _, val := range usbArr {
		computer.Working(val)
	}
}