package main

import (
	controller "github.com/catdog93/GoIT/controller"
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func CreateUrlMappings() {
	router = gin.Default()
	group := router.Group(string(controller.RelativePathEmployees))
	{
		group.GET(string(controller.ReadAllEmployees), controller.ReadAll)
		group.POST(string(controller.ReadAllEmployees), controller.ReadAllPost)
		group.GET(string(controller.ReadEmployee), controller.ReadId)
		group.POST(string(controller.CreateEmployee), controller.Create)
		group.PUT(string(controller.ReplaceEmployee), controller.ReplaceId)
		group.DELETE(string(controller.DeleteEmployee), controller.DeleteId)

	}
}

func main() {
	// gin web API
	CreateUrlMappings()
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
