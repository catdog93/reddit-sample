package main

import (
	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"testTaskBitmediaLabs/controller"
	"testTaskBitmediaLabs/service"

	//"github.com/go-redis/redis/v7"
	"log"
)

var (
	router = gin.Default()
)

func main() {
	ConfigureDBConnection()

	ConfigureRouter()

	err := router.Run() // listen and serve on localhost:8080
	if err != nil {
		log.Fatal(err)
	}
}

func ConfigureDBConnection() {
	session, err := mgo.Dial(service.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	service.PostsCollection = session.DB(service.DBName).C(service.PostsCollectionName)
	service.UsersCollection = session.DB(service.DBName).C(service.UsersCollectionName)
	ai.Connect(session.DB(service.DBName).C(service.PostsCollectionName))
	ai.Connect(session.DB(service.DBName).C(service.UsersCollectionName))
}

func ConfigureRouter() {
	InitMblogRouterGroup()

	InitUsersRouterGroup()

	router.LoadHTMLGlob("../templates/*.html")
	router.Static("files", "../templates")
}

func InitUsersRouterGroup() {
	usersGroup := router.Group(controller.Users)
	usersGroup.Use(controller.CheckIsAuthorised)

	usersGroup.POST(controller.SignupURI, controller.SignupPost)
	usersGroup.GET(controller.SignupURI, controller.GetSignupForm)
	usersGroup.POST(controller.SigninURI, controller.SigninPost)
	usersGroup.GET(controller.SigninURI, controller.GetSigninForm)
}

func InitMblogRouterGroup() {
	mblogGroup := router.Group(controller.MblogURI)
	mblogGroup.Use(controller.TokenAuth)

	mblogGroup.GET(controller.HomeURI, controller.GetHomePage)
	mblogGroup.GET(controller.CreatePostURI, controller.GetCreatePostForm)
	mblogGroup.POST(controller.CreatePostURI, controller.CreatePost)
	mblogGroup.POST(controller.SubscriptionsURI, controller.SubscribeUser)
	mblogGroup.POST(controller.UnfollowURI, controller.UnfollowUser)
	mblogGroup.GET(controller.SubscriptionsURI, controller.GetSubscriptionsPage)
}
