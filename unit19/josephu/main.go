package main

import (
	"fmt"
)

type Boy struct {
	No int
	Next *Boy
}

// num -> boy 个数
// *Boy -> 该环形链表第一个 Boy 的指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num too low......")
		return first
	}
	// 构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy {
			No : i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first
		}
	}

	return first
}

func ShowBoy(first *Boy) {
	if first.Next == nil {
		fmt.Println("empty......")
		return
	}
	
	curBoy := first
	for {
		fmt.Printf("Boy[%d] => ", curBoy.No)
		if curBoy.Next == first {
			break
		}
		curBoy = curBoy.Next
	}
	fmt.Println()
}

// 1、编写一个函数 PlayGame(first *Boy, startNo int, countNo int)
// 2、按照要求，在环形链表中留下最后一个
func PlayGame(first *Boy, startNo int, countNo int) {
	if first.Next == nil {
		fmt.Println("empty......")
		return
	}
	// 3、startNo 校验
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}
	// 4、让 first 移动到 startNo - 1，后面删除谁由 first 为准
	for i := 1; i <= startNo - 1; i++ {
		first = first.Next
		tail = tail.Next
	}
	// 5、开始数 countNo - 1，然后删除 first 指向的 Boy
	for {
		for i := 1; i <= countNo - 1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Println(first.No, "bye bye......")
		first = first.Next
		tail.Next = first
		// 退出条件 -> 只剩一个
		if tail == first {
			break
		}
	}
	fmt.Println(first.No, "win......")
}

func main()  {
	first := AddBoy(5)
	ShowBoy(first)
	PlayGame(first, 2, 3)
}