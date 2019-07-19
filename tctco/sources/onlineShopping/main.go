package main

import (
	"fmt"
	"onlineShopping/controller"
	"onlineShopping/model"
	"onlineShopping/view"
)

func main() {
	err := model.DBInit()
	if err != nil {
		fmt.Println("不能创建数据库池！")
		return
	}

	authMiddleWare := controller.MiddleWareInit()

	r := view.RouterInit(authMiddleWare)
	r.Run(":8080")
}
