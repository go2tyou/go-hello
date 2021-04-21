package main

import (
	_"fmt"
)

type B interface {
	testB()
}

type C interface {
	testC()
}

type A interface {
	B
	C
	testA()
}

type Stu struct {

}

func (stu Stu) testA() {
	
}

func (stu Stu) testB() {
	
}

func (stu Stu) testC() {
	
}

func main()  {
	var stu Stu
	var a A = stu
	a.testA()
}