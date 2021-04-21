package main

import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (stu *Student) say() string {
	info := fmt.Sprintf("stu name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return info
}

type Box struct {
	len float64
	wid float64
	height float64
}

func (box *Box) getVolum() float64 {
	return box.len * box.wid * box.height
}

type Visitor struct {
	name string
	age int
}

func (visitor *Visitor) showPrice() {
	if visitor.age >= 90 || visitor.age <= 8 {
		fmt.Println("age no...")
		return
	}
	if visitor.age > 18 {
		fmt.Printf("name=%v age=%v price=20\n", visitor.name, visitor.age)
	} else {
		fmt.Printf("name=%v age=%v price=0\n", visitor.name, visitor.age)
	}
}

func main()  {
	var stu = Student {
		name : "tom",
		gender : "male",
		age : 17,
		id : 100,
		score : 99,
	}
	fmt.Println(stu.say()) // stu name=[tom] gender=[male] age=[17] id=[100] score=[99]

	var box Box
	box.len = 1.2
	box.wid = 2.3
	box.height = 3.4
	v := box.getVolum()
	fmt.Printf("v=%.2f\n", v) // v=9.38

	var visitor Visitor
	for {
		fmt.Println("input name:")
		fmt.Scanln(&visitor.name)
		if visitor.name == "n" {
			fmt.Println("exit...")
			break
		}
		fmt.Println("input age:")
		fmt.Scanln(&visitor.age)
		visitor.showPrice()
	}
}