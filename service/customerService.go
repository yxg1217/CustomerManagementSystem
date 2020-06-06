package service

import (
	"CustomerManagementSystem/model"
)

//The CustomerService completes CRUD operations on customers 该CustomerService,完成对客户的操作,包括增删改

type CustomerService struct {
	customers []model.Customer
	//Declare a field indicating how many customers the current slice contains
// Behind this field, it can also be used as the id+1 for new customers
	customerNum int
}

//Write a method that can return *CustomerService
func NewCustomerService() *CustomerService {
	//To be able to see the customer in the slice, we initialize a customer
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "Jean", "Homme", 20, "+33 613749302", "jean@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//Return customer slice
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//Add customers to customers slice
func (this *CustomerService) Add(customer model.Customer) bool {

	//Use the added order as the Id
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//Delete customer by id (delete from slice)
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//If index == -1, there is no such client
	if index == -1 {
		return false 
	}
	//How to remove an element from a slice
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//Find the corresponding subscript in the slice according to the id, if there is no such customer, return -1
func (this *CustomerService) FindById(id int)  int {
	index := -1
	//Iterate through this.customers slice 遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//found it
			index = i
		}
	}
	return index
}