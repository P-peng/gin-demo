package router

import (
	. "gin-demo/api/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users", Users)

	router.GET("/getUser", GetUser)

	router.GET("/insert", Insert)

	router.GET("/update", Update)

	router.GET("/delete", Delete)

	return router
}
