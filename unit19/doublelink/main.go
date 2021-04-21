package main

import (
	"fmt"
)

type HeroNode struct {
	no int
	name string
	pre *HeroNode // 指向前一个结点
	next *HeroNode // 指向下一个结点
}

// 插入一个结点
// 方式一：单链表最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 1、先找到该链表的最后一个结点
	// 2、创建一个辅助结点
	temp := head
	for {
		if temp.next == nil { // 找到了最后
			break
		}
		temp = temp.next // 让 temp 指向下一个结点
	}
	// 3、将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 方式二：根据 no 从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 1、找到适当的结点
	temp := head
	// 让插入结点 no 和 temp 的下一个结点的 no 比较
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no {
			// 说明 newHeroNode 就应该插入到 temp 后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明链表中已经存在这个 no 不允许插入
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("此人已在......", newHeroNode.no)
		return
	}
	newHeroNode.next = temp.next
	newHeroNode.pre = temp
	if temp.next != nil {
		temp.next.pre = newHeroNode
	}
	temp.next = newHeroNode
}

func DelHeroNode(head *HeroNode, no int) {
	temp := head
	// 1、找到要删除的结点
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == no {
			// 说明找到
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	} else {
		fmt.Println("要删除的结点不存在......", no)
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空的......")
		return
	} 
	for {
		fmt.Printf("[%d, %s] => ", temp.next.no, temp.next.name)
		temp = temp.next
		// 判断是否到链表最后
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空的......")
		return
	} 
	// 让 temp 定位到最后一个结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	for {
		fmt.Printf("[%d, %s] => ", temp.no, temp.name)
		temp = temp.pre
		// 判断是否到链表头
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

func main()  {
	// 1、创建一个头结点
	head := &HeroNode{}
	// 2、创建一个新的结点
	hero1 := &HeroNode {
		no : 1,
		name : "aaa",
	}
	hero2 := &HeroNode {
		no : 2,
		name : "bbb",
	}
	hero3 := &HeroNode {
		no : 3,
		name : "ccc",
	}
	hero4 := &HeroNode {
		no : 4,
		name : "ddd",
	}
	// 3、插入
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero4)
	// 4、显示
	ListHeroNode(head)
	ListHeroNode2(head)
	DelHeroNode(head, 4)
	ListHeroNode(head)
}