package service

import (
	"fmt"
	"go_batch_create_orders/model"
	"strconv"
)

type UserService struct {
	Db
}

func (m *UserService) AddUser(user *model.Users) (res int64) {
	insertData := make([]model.Users, 0, 10)
	for j := 0; j < 10; {
		insertData = append(insertData, model.Users{UserName: "用户" + strconv.Itoa(j) + "号"})
		j++
	}
	a := new(Add)
	a.Resouces = insertData
	a.Table = "users"
	StartTransaction()
	res, err := a.Insert()
	if err != nil {
		fmt.Println(err.Error())
		res = 0
	}
	CommitTransaction()
	return
}

func (m *UserService) DeleteUser(id int) {
	StartTransaction()
	d := Delete{Where: "id = " + strconv.Itoa(id), Table: "users", Resouces: new(model.Users), Unscoped: false}
	res, err2 := d.Delete()
	RollbackTransaction()
	fmt.Println(res)
	fmt.Println(err2)
}

func (m *UserService) UpdateUser() {
	u := new(Update)
	u.Table = "users"
	u.Resouces = model.Users{UserName: "修改的用户1"}
	u.Where = "id = 3"
	StartTransaction()
	res, err := u.Update()
	CommitTransaction()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
