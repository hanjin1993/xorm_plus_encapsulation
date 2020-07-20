package model

//`id``order_sn``user_id``create_time``status``address`
type Order struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrderSn    string `json:"order_sn" xorm:"not null CHAR(30)"`
	UserId     int    `json:"user_id" xorm:"not null int(11) index"`
	CreateTime int    `json:"create_time" xorm:"not null int(10) created"`
	Status     int8   `json:"status" xorm:"not null tinyint(1)"`
	Address    string `json:"address" xorm:"not null varchar(255)"`
}

type OrderGoods struct {
	Id         int `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrderId    int `json:"order_id" xorm:"not null int(11) index"`
	GoodsId    int `json:"goods_id" xorm:"not null int(11) index"`
	Num        int `json:"num" xorm:"not null int(11)"`
	CreateTime int `json:"create_time" xorm:"not null int(10) created"`
}

type UserOrderGoodsList struct {
	Order      `xorm:"extends"`
	OrderGoods `xorm:"extends"`
	Users      `xorm:"extends"`
}
