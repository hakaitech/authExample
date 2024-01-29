package main

import (
	"authExample/apis"
	"authExample/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/helloworld", apis.HelloWord)
	r.GET("/gt", apis.GenToken)
	r.Use(middlewares.Validate)
	r.GET("/protected/hw", apis.HelloWord)
	r.Run(":8080")
}
