package service

import (
	"fmt"
	"go_batch_create_orders/model"
	"strconv"
)

type OrderService struct {
	Db
}

func (OrderService) GetOrderList() {
	var res []model.UserOrderGoodsList
	res = make([]model.UserOrderGoodsList, 0)
	join := make([][4]string, 0, 3)
	join = append(join, [4]string{"left", "order_goods", "OG", "OG.order_id = O.id"})
	join = append(join, [4]string{"inner", "users", "U", "U.id = O.user_id"})
	where := "O.order_sn='OR1234567'"
	query := InitQuery()
	query.Resouces = &res
	query.Where = where
	query.Join = join
	query.Alias = "O"
	query.Table = "order"
	query.Page = 1
	query.PageSize = 2
	query.OrderBy = "O.id desc"
	query.GroupBy = ""
	err := query.Select()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
	for _, v := range res {
		fmt.Println(v.OrderGoods)
	}
}

func (OrderService) InsertData() {
	params := make([]model.Users, 0, 100)
	j := 0
	for j <= 10 {
		params = append(params, model.Users{UserName: "test" + strconv.Itoa(j)})
		j++
	}
	add := Add{Resouces: params, Table: "users"}
	insert, err := add.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(insert)
}
