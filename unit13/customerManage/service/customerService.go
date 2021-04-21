package service

import (
	"gocode/helloworld/unit13/customerManage/model"
)

type CustomerService struct {
	customers []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "草爹", "famale", 18, "123456", "123456@163.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (this *CustomerService) DeleteById(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	// [0, index)
	this.customers = append(this.customers[:index], this.customers[index + 1:]...)
	return true
}