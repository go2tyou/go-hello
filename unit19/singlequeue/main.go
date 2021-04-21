package main

import (
	"fmt"
	"errors"
	"os"
)

type Queue struct {
	maxSize int
	array [5]int // 数组模拟队列
	front int // 队首 不包含队首
	rear int // 队尾 包含队尾
}

func (this *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if this.rear == this.maxSize - 1 { // rear 包含最后一个元素
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

// 找到队首，然后遍历到队尾
func (this *Queue) ShowQueue() {
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, this.array[i])
	}
}

func main()  {
	queue := &Queue {
		maxSize : 5,
		front : -1,
		rear : -1,
	}
	
	var key string
	var val int
	for {
		fmt.Println("输入 add  添加：")
		fmt.Println("输入 get  获取：")
		fmt.Println("输入 show 显示：")
		fmt.Println("输入 exit 退出：")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("添加。。。。。。")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("ok")
			}
		case "get":
			fmt.Println("获取。。。。。。")
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("ok ->", val)
			}
		case "show":
			fmt.Println("显示。。。。。。")
			queue.ShowQueue()
		case "exit":
			fmt.Println("退出。。。。。。")
			os.Exit(0)
		}
	}
}