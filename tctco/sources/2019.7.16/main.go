package main

import (
	"demo/controller"
	"demo/view"
)

func main() {
	authMiddleware := controller.MiddleWareInit()   // this middle ware cannot tell the authority
	r := view.RouterInit(authMiddleware) // authority is evaluated by the handlers!

	r.Run(":8080")
}