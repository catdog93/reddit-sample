package controller

import (
	"fmt"
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
		_, ok := service.TokensCache[token]
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
	user, _ := service.TokensCache[token]
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
	emailOfPublisher := entity.Email{}
	err := context.BindJSON(&emailOfPublisher)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	subscriberUser, ok := service.TokensCache[token]
	if ok {
		subscriber, err := service.FindUserByEmail(subscriberUser.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		wasAlreadySubscribed, err := service.Subscribe(subscriber, emailOfPublisher.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		if !wasAlreadySubscribed {
			//context.String(http.StatusOK,"Successfully subscribed %s", emailOfPublisher.Email)
			context.JSON(http.StatusOK, entity.Message{Message: fmt.Sprintf("Successfully subscribed %s", emailOfPublisher.Email)})
		}
	}
}

func UnfollowUser(context *gin.Context) {
	emailOfPublisher := entity.Email{}
	err := context.BindJSON(&emailOfPublisher)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	subscriberUser, ok := service.TokensCache[token]
	if ok {
		subscriber, err := service.FindUserByEmail(subscriberUser.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		wasNotFollowed, err := service.Unfollow(subscriber, emailOfPublisher.Email)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		if !wasNotFollowed {
			//context.String(http.StatusOK,"Successfully unfollowed from %s", emailOfPublisher.Email)
			context.JSON(http.StatusOK, entity.Message{Message: fmt.Sprintf("Successfully unfollowed from %s", emailOfPublisher.Email)})
		}
	}
}
