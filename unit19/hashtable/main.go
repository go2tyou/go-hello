package main

import (
	"fmt"
)

type Emp struct {
	Id int
	Name string
	Next *Emp
}

func (this *Emp) Show() {
	fmt.Printf("链表 %d => emp[id=%d name=%s]\t", this.Id % 7, this.Id, this.Name)
}

// 不带表头
type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	// 定义两个辅助指针
	var pre *Emp
	cur := this.Head
	if cur == nil {
		this.Head = emp
		return
	}
	// 找到 emp 合适的位置
	// 让 cur 和 emp 比较，然后让 pre 保持在 cur 前面
	for {
		if cur != nil {
			if cur.Id >= emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	// 退出时
	pre.Next = emp
	emp.Next = cur
}

func (this *EmpLink) ShowLink(linkno int) {
	if this.Head == nil {
		fmt.Println("链表为空......", linkno)
		return
	}
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表 %d => emp[id=%d name=%s]\t", linkno, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) ListAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func (this *HashTable) FindById(id int) {
	linkno := this.HashFun(id)
	emp := this.LinkArr[linkno].FindById(id)
	if emp == nil {
		fmt.Println(id, "不存在......")
	} else {
		emp.Show()
	}
}

func (this *HashTable) Insert(emp *Emp) {
	// 使用散列函数确定位置
	linkno := this.HashFun(emp.Id)
	// 使用对应的链表添加
	this.LinkArr[linkno].Insert(emp)
}

func (this *HashTable) HashFun(id int) int {
	return id % 7 // 对应链表的下标
}

func main()  {
	emp1 := &Emp{
		Id : 1,
		Name : "s",
	}
	emp2 := &Emp{
		Id : 2,
		Name : "ss",
	}
	emp3 := &Emp{
		Id : 3,
		Name : "sss",
	}
	emp23 := &Emp{
		Id : 23,
		Name : "dd",
	}
	emp47 := &Emp{
		Id : 47,
		Name : "d",
	}
	var hashTable HashTable
	// hashTable.ListAll()
	hashTable.Insert(emp1)
	hashTable.Insert(emp3)
	hashTable.Insert(emp2)
	hashTable.Insert(emp23)
	hashTable.Insert(emp47)
	fmt.Println("list start......")
	hashTable.ListAll()
	fmt.Println("list end......")
	fmt.Println()
	hashTable.FindById(4)
	hashTable.FindById(3)
}