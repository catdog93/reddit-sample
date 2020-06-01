package main

import (
	"fmt"
	controller "github.com/catdog93/GoIT/employeeController"
	"github.com/catdog93/GoIT/repository"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var router *gin.Engine

func CreateUrlMappings() {
	router = gin.Default()
	group := router.Group(string(controller.RelativePathEmployees))
	{
		//group.POST(string(controller.ReadAllEmployees), controller.ReadAll)
		group.POST(string(controller.CreateEmployee), controller.Create)
		group.GET(string(controller.ReadEmployee), controller.ReadId)
		group.POST(string(controller.ReadEmployee), controller.ReadIdPost)
		//group.POST("/login/", controller.Login)
		//group.PUT("/users/:id", controller.UpdateUser)
		//group.POST("/users", controller.PostUser)
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

func func1() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST", "https://google.com", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func f1() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://google.com", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func f2() {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func f() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		c.HTML(200, "helloUser.html", repository.Obj{"name": name})
	})
	r.Run() // listen and serve on localhost:8080 (for windows "localhost:8080")
}
