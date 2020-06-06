package model

import (
	"fmt"
)

//Declare a customer structure, indicating a customer information

type Customer struct {
	Id int
	Name string
	Sex string
	Age int
	Phone string
	Email string
}

//Write a factory pattern that returns an instance of a customer
//编写一个工厂模式,返回一个客户的实例
func NewCustomer(id int, name string, sex string,
	age int, phone string, email string ) Customer {
		return Customer{
			Id: id,
			Name: name,
			Sex: sex,
			Age: age,
			Phone: phone,
			Email: email,
		}
	}

//Method for creating customer instance without Id 不带Id的创建客户实例方法
func NewCustomer2(name string, sex string,
	age int, phone string, email string ) Customer {
		return Customer{
			//Id: id,
			Name: name,
			Sex: sex,
			Age: age,
			Phone: phone,
			Email: email,
		}
	}

	//Return user information
	func (this *Customer) GetInfo() string {
		info := fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, 
			this.Sex, this.Age, this.Phone, this.Email)
		return info
	}