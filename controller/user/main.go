package user

import (
	"github.com/gin-gonic/gin"
	"go_batch_create_orders/model"
	"go_batch_create_orders/service"
	"go_batch_create_orders/utils"
)

func AddUser(ctx *gin.Context) {
	user := new(model.Users)
	err := ctx.ShouldBind(&user)
	u := new(utils.JsonData)
	if err != nil {
		u.Code = 502
		u.Message = "添加失败"
		ctx.JSON(502, u)
	}
	s := new(service.UserService)
	res := s.AddUser(user)
	if res > 0 {
		u.Code = 200
		u.Message = "添加成功"
		ctx.JSON(200, u)
	} else {
		u.Code = 502
		u.Message = "添加失败"
		ctx.JSON(502, u)
	}
}

func DeleteUser(ctx *gin.Context) {
	s := new(service.UserService)
	s.DeleteUser(3)
}

func UpdateUser(ctx *gin.Context) {
	new(service.UserService).UpdateUser()

}
