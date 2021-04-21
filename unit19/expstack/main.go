package main

import (
	"fmt"
	"errors"
	"strconv"
)

// 使用数组模拟一个栈
type Stack struct {
	MaxTop int // 最大可存放数
	Top int // 栈顶
	// 栈底固定
	arr [20]int // 数组模拟栈
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

// 判断一个字符是不是一个运算符
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

// 运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误......")	
	}
	return res
}

// 返回某个运算符的优先级
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main()  {
	numStack := &Stack {
		MaxTop : 20,
		Top : -1,
	}
	operStack := &Stack {
		MaxTop : 20,
		Top : -1,
	}
	exp := "300+2*16-21"
	// 辅助扫描 exp
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for {
		// 处理多位数 ?
		ch := exp[index:index + 1]
		// 字符对应的 ASCII 值
		temp := int([]byte(ch)[0])
		// 判断是否为符号
		if operStack.IsOper(temp) {
			// 空栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			}			
			// 优先级判断
			if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
				num1, _ = numStack.Pop()
				num2, _ = numStack.Pop()
				oper, _ = operStack.Pop()
				result = operStack.Cal(num1, num2, oper)
				// 将计算结果重新入数栈
				numStack.Push(result)
				// 将当前符号入符号栈
				operStack.Push(temp)
			} else {
				operStack.Push(temp)
			}
		} else {
			// 处理多位数
			// 1、定义一个变量 keepNum string 做拼接
			keepNum += ch
			// 2、每次向 index 的后面测试一下，看是否为运算符，然后再做处理
			if index == len(exp) - 1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向 index 的后面探测
				if operStack.IsOper(int([]byte(exp[index+1:index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}
		// 继续扫描
		if index == len(exp) - 1 {
			break
		} else {
			index++
		}
	}

	// 表达式扫描完毕后
	for {
		if operStack.Top == -1 {
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		// 将计算结果重新入数栈
		numStack.Push(result)
	}
	result, _ = numStack.Pop()
	fmt.Println("exp=", exp)
	fmt.Println("result=", result)
}