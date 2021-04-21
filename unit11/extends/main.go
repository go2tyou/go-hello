package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("name=%v age=%v score=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student // 嵌入了 Student 的匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("p testing...")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("g testing...")
}

// 继承
// 代码复用性、维护性、扩展性提高
func main()  {
	// 结构体嵌入匿名结构体后使用
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 10
	pupil.testing()
	pupil.Student.SetScore(99)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "tom"
	graduate.Student.Age = 10
	graduate.testing()
	graduate.Student.SetScore(59)
	graduate.Student.ShowInfo()
}