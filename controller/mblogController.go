package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
)

const (
	MblogURI         = "/mblog"
	HomeURI          = "/home"
	SubscriptionsURI = "/subscriptions"
	UnfollowURI      = "/unfollow"
)

func TokenAuth(context *gin.Context) {
	token, err := context.Cookie(tokenString)
	if err == nil {
		_, ok := service.Tokens[token]
		if ok {
			return
		}
	}
	context.Redirect(http.StatusMovedPermanently, Users+SigninURI)
	context.Abort()
}

func GetHomePage(context *gin.Context) {
	posts, err := service.GetPostsWithEmail()
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.HTML(http.StatusOK, "home.html", posts)
}

func GetSubscriptionsPage(context *gin.Context) {
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, _ := service.Tokens[token]
	userFromDB, err := service.FindUserByEmail(user.Email)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	posts, err := service.GetPostsWithEmail(userFromDB.UsersIDs...)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	context.HTML(http.StatusOK, "subscriptions.html", posts)
}

func SubscribeUser(context *gin.Context) {
	email := entity.Email{}
	err := context.BindJSON(&email)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, ok := service.Tokens[token]
	if ok {
		userFromDB, err := service.FindUserByEmail(user.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		err = service.Subscribe(userFromDB, email.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func UnfollowUser(context *gin.Context) {
	email := entity.Email{}
	err := context.BindJSON(&email)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, ok := service.Tokens[token]
	if ok {
		userFromDB, err := service.FindUserByEmail(user.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		err = service.Unfollow(userFromDB, email.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}
