package model

type Users struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserName   string `json:"user_name" xorm:"CHAR(30) not null " binding:"required"`
	Password   string `json:"password" xorm:"VARCHAR(50) not null" binding:"required"`
	CreateTime int    `json:"create_time" xorm:"INT(11) created"`
	DeleteTime int    `json:"delete_time" xorm:"INT(10) deleted"`
	Sex        int8   `json:"sex" xorm:"TINYINT(1) default 2"`
}
