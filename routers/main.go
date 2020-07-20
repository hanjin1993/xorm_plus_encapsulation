package routers

import (
	"github.com/gin-gonic/gin"
	"go_batch_create_orders/controller/order"
	"go_batch_create_orders/controller/user"
	"go_batch_create_orders/utils"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			data := utils.JsonData{Code: 501, Message: "暂未登录"}
			context.JSON(http.StatusNetworkAuthenticationRequired, data.New())
			context.Abort()
		}
		if token == "qwertyuiop122321" {
			context.Next()
		}
	}
}

func RouterStart() *gin.Engine {
	//创建路由
	engine := gin.Default()
	engine.Use(Middleware())
	//配置路由等策略
	engine.Group("/")
	{
		engine.POST("batch_order", order.GetOrderList)
		engine.POST("add_order", order.AddOrder)
		engine.POST("add_user", user.AddUser)
		engine.POST("del_user", user.DeleteUser)
		engine.POST("update_user", user.UpdateUser)
	}
	return engine
}
