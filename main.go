package main

import (
	"go_batch_create_orders/routers"
	"log"
)

func main() {
	router := routers.RouterStart()
	err := router.Run(":8088")
	if err != nil {
		log.Println("服务器启动失败")
	}
}
