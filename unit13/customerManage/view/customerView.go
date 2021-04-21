package main

import (
	"fmt"
	"gocode/helloworld/unit13/customerManage/service"
	"gocode/helloworld/unit13/customerManage/model"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("```````````````````````客户列表`````````````````````````````")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("```````````````````````客户列表完成`````````````````````````````\n\n")
}

func (this *customerView) add() {
	fmt.Println("```````````````````````添加客户`````````````````````````````")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Printf("```````````````````````添加客户完成`````````````````````````````\n\n")
	} else {
		fmt.Printf("```````````````````````添加客户失败`````````````````````````````\n\n")
	}
}

func (this *customerView) delete() {
	fmt.Println("```````````````````````删除客户`````````````````````````````")
	fmt.Println("请删除待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.DeleteById(id) {
			fmt.Printf("```````````````````````删除客户完成`````````````````````````````\n\n")
		} else {
			fmt.Printf("```````````````````````删除客户失败，客户编号不存在```````````````\n\n")
		}
	}
}

func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}
		fmt.Println("输入有误，确认是否退出(Y/N)：")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	}
}

func (this *customerView) mainView() {
	for {
		fmt.Println("``````````````````````客户信息管理软件``````````````````````````")
		fmt.Println("```````````````````````1、添加客户`````````````````````````````")
		fmt.Println("```````````````````````2、修改客户`````````````````````````````")
		fmt.Println("```````````````````````3、删除客户`````````````````````````````")
		fmt.Println("```````````````````````4、客户列表`````````````````````````````")
		fmt.Println("```````````````````````5、退   出``````````````````````````````")
		fmt.Println("请选择(1~5)：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("输入有误，重新输入：")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("您退出了客户信息管理软件！")
}

func main()  {
	customerView := customerView {
		key : "",
		loop : true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainView()
}