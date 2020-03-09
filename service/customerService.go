package service

import (
	"CustomerManagementSystem/model"
)

//该CustomerService,完成对客户的操作,包括增删改

type CustomerService struct {
	customers []model.Customer
	//声明一个字段,表示当前切片含有多少个客户
	//该字段后面,还可以作为新客户的id+1
	customerNum int
}

//编写一个方法,可以反悔 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到客户在切片中,我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "Jean", "Homme", 20, "+33 613749302", "jean@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户到customers切片
func (this *CustomerService) Add(customer Customer) bool {

	//用添加的顺序作为Id
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}