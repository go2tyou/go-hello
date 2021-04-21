package main

import (
	"fmt"
	"errors"
	"os"
)

type CirCleQueue struct {
	maxSize int
	array [5]int
	head int // 队首 包含队列队首元素
	tail int // 队尾 不包含队列最后一个元素
}

func (this *CirCleQueue) AddQueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

func (this *CirCleQueue) GetQueue() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, err
}

func (this *CirCleQueue) ShowQueue() {
	// 判断 size
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空。。。。。。。")
	}
	fmt.Println("队列如下。。。。。。。")
	temp := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", temp, this.array[temp])
		temp = (temp + 1) % this.maxSize
	}
}

// 判断是否满
func (this *CirCleQueue) IsFull() bool {
	return (this.tail + 1) % this.maxSize == this.head
}

// 判断是否空
func (this *CirCleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 元素个数
func (this *CirCleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main()  {
	queue := &CirCleQueue {
		maxSize : 5,
		head : 0,
		tail : 0,
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