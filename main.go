package main

import (
	"github.com/Ivanhahanov/LDAPManager/routers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.Use(cors.Default())
	router.POST("/user/add/", routers.AddUserHandler)
	router.GET("/user/:username/", routers.GetUserHandler)
	router.Run(":8081")
}
