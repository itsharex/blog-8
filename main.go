package main

import (
	"boKe/model"
	"boKe/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
