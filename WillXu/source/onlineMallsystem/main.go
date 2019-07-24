package main

import "onlineMallsystem/routers"

func main() {
	r := routers.InitRouter()
	_ = r.Run(":8080")
}
