package order

import (
	"github.com/gin-gonic/gin"
	"go_batch_create_orders/service"
)

func AddOrder(ctx *gin.Context) {
	//首先对数据进行过滤

	new(service.OrderService).InsertData()
	//header := ctx.GetHeader("token")
	//mysql.New()
	//ctx.JSON(http.StatusOK, utils.JsonData{Code: 200, Message: "成功了", Data: header})
}

func GetOrderList(ctx *gin.Context) {
	//首先对数据进行过滤
	new(service.OrderService).GetOrderList()
	//header := ctx.GetHeader("token")
	//mysql.New()
	//ctx.JSON(http.StatusOK, utils.JsonData{Code: 200, Message: "成功了", Data: header})
}
