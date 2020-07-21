package main

import (
	_ "gin-demo/api/database"
	orm "gin-demo/api/database"
	"gin-demo/api/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}
