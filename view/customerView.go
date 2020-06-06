package main

import (
	"fmt"
	"CustomerManagementSystem/service"
	"CustomerManagementSystem/model"
)

type customerView struct {

	//Define the necessary fields
	key string //Receive user input
	loop bool // Main menu indicating whether to cycle 表示是否循环的主菜单
	//Add a field customerService, otherwise customerService cannot be called
	customerService *service.CustomerService
}

//Display all customer information
func (this *customerView) list() {

	//获取所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("-------------------- Customer List --------------------")
	fmt.Println("ID\tName\tGender\tAge\tPhone\tEmail")
	for i := 0; i < len(customers); i++ {
		// fmt.Println(customers[i].Id, "\t", customers[i].Name,) //这样写不如在customer里封装一个方法
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------------------------- End -------------------------")
}

//得到用户的输入,信息构建新的客户,并完成添加
func (this *customerView) add() {
	fmt.Println("--------------------- Add customer --------------------")
	fmt.Println("Name:")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("Sex:")
	sex := ""
	fmt.Scanln(&sex)

	fmt.Println("Age:")
	age := ""
	fmt.Scanln(&age)

	fmt.Println("Phone:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("Email:")
	email := ""
	fmt.Scanln(&email)

	//构建一个新的customer实例
	//id号是唯一的,需要系统分配
	customer := model.NewCustomer2(name, sex, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-------------------Add Customer Successfully--------------------")
	} else {
		fmt.Println("----------------------Failed Add Customer-----------------------")
	}
}
//Show main menu
func (this *customerView) mainMenu() {

	for {
		fmt.Println("-------------------- Customer Management System --------------------")
		fmt.Println("                     1 Add customer")
		fmt.Println("                     2 Modify customer")
		fmt.Println("                     3 Delete customer")
		fmt.Println("                     4 Customer List")
		fmt.Println("                     5 Exit")
		fmt.Println("Please enter your order number (1-5): ")

		fmt.Scanln(&this.key)

		switch this.key {
			case "1" :
				this.add()
			case "2" :
				fmt.Println("Modify customer")
			case "3" :
				fmt.Println("Delete customer")
			case "4" :
				this.list()
			case "5" :
				this.loop = false
			default : 
				fmt.Println("Oups! Please enter the correct option!")
		}

		if !this.loop {
			break
		}

	}

	fmt.Println("Bye, see you next time ^_^ !")
}

func main() {

	//在主函数中创建一个customerView实例,并显示主菜单
	customerView := customerView{
		key : "",
		loop : true,
	}

	//在这里完成对customerView结构体的customerService字段的初始化工作
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()

}