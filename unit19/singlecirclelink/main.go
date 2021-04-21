package main

import (
	"fmt"
)

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 判断添加的是否为第一个结点
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 形成一个环状
		fmt.Println(newCatNode, "加入环形链表......")
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到环形链表
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCatNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空......")
		return
	}
	for {
		fmt.Printf("cat[%d] = %s ---> ", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	fmt.Println()
}

func DelCatNode(head *CatNode, no int) *CatNode {
	temp := head
	helper := head
	if temp.next == nil {
		fmt.Println("空，不删......")
		return head
	}
	// 只有一个结点
	if temp.next == head {
		temp.next = nil
		return head
	}

	// 将 helper 定位到环形链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		// 转了一圈，但还没比较最后一个
		if temp.next == head {
			break
		}
		if temp.no == no {
			// 找到
			if temp == head { // 删除的是头结点
				head = head.next
			}
			helper.next = temp.next
			fmt.Println("cat del......", no)
			flag = false
			break
		}
		temp = temp.next // 用于比较
		helper = helper.next // 用于辅助删除
	}

	if flag {
		if temp.no == no {
			helper.next = temp.next
			fmt.Println("cat del......", no)
			flag = false
		} else {
			fmt.Println("cat not found......", no)
		}
	}

	return head
}

func main()  {
	head := &CatNode{}

	cat1 := &CatNode {
		no : 1,
		name : "r",
	}
	cat2 := &CatNode {
		no : 2,
		name : "rr",
	}
	cat3 := &CatNode {
		no : 3,
		name : "rrr",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCatNode(head)
	DelCatNode(head, 2)
	DelCatNode(head, 4)
	ListCatNode(head)
}