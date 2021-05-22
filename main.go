package main

import (
	_ "gin-app/initDB"
	"gin-app/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()
}
