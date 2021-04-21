package main

import (
	"fmt"
	"errors"
)

// 使用数组模拟一个栈
type Stack struct {
	MaxTop int // 最大可存放数
	Top int // 栈顶
	// 栈底固定
	arr [5]int // 数组模拟栈
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	// 先判断栈是否满
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full......")
		return errors.New("stack full")
	}
	this.Top++
	// 放入数据
	this.arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty......")
		return 0, errors.New("stack empty")
	}
	// 先取值，再 Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈，需要从栈顶开始遍历
func (this *Stack) List() {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty......")
		return
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])		
	}
}

func main()  {
	stack := &Stack {
		MaxTop : 5,
		Top : -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	// stack.Push(6) // stack full......
	stack.List()
	fmt.Println()
	
	fmt.Println("stack pop start......")
	val, _ := stack.Pop()
	fmt.Println("pop val=", val)
	stack.List()
	fmt.Println()

	val, _ = stack.Pop()
	fmt.Println("pop val=", val)
	stack.List()
	fmt.Println()

	val, _ = stack.Pop()
	fmt.Println("pop val=", val)
	stack.List()
	fmt.Println()

	val, _ = stack.Pop()
	fmt.Println("pop val=", val)
	stack.List()
	fmt.Println()

	val, _ = stack.Pop()
	fmt.Println("pop val=", val)
	stack.List()
	fmt.Println()

	val, _ = stack.Pop()
	fmt.Println("pop val=", val)
	stack.List()
	fmt.Println()
}