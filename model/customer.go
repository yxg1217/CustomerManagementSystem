package model

import (
	"fmt"
)

//声明一个客户结构体,表示一个客户信息

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

//编写一个工厂模式,返回一个客户的实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string ) Customer {
		return Customer{
			Id: id,
			Name: name,
			Gender: gender,
			Age: age,
			Phone: phone,
			Email: email,
		}
	}

	//返回用户的信息
	func (this *Customer) GetInfo() string {
		info := fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, 
			this.Gender, this.Age, this.Phone, this.Email)
		return info
	}