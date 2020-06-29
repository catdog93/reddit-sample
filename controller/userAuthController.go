package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
)

const (
	Users     = "/users"
	SigninURI = "/signin"
	SignupURI = "/signup"

	incorrectEmailOrPassword = "Incorrect email or password"
	suchEmailAlreadyExists   = "Such email already exists"

	email    = "inputEmail"
	password = "inputPassword"
)

// redirect user to home page if he have already authorized
func CheckIsAuthorised(context *gin.Context) {
	token, err := context.Cookie(tokenString)
	if err != nil {
		return
	}
	_, ok := service.Tokens[token]
	if ok {
		context.Redirect(http.StatusMovedPermanently, MblogURI+HomeURI)
		context.Abort()
	}
}

func SignupPost(context *gin.Context) {
	email := context.PostForm(email)
	pass := context.PostForm(password)

	user := entity.User{
		Email:    email,
		Password: pass,
	}

	err := service.CreateUser(user)
	if err != nil {
		context.HTML(http.StatusOK, "signup.html", suchEmailAlreadyExists)
		return
	}
	context.Redirect(http.StatusMovedPermanently, Users+SigninURI)
}

func SigninPost(context *gin.Context) {
	email := context.PostForm(email)
	password := context.PostForm(password)

	userBody := entity.UserBody{
		Email:    email,
		Password: password,
	}
	user, err := service.FindUserByEmail(userBody.Email)
	if err != nil && user != nil {
		context.HTML(http.StatusOK, "signin.html", incorrectEmailOrPassword)
		return
	}
	token := service.CreateToken(user)
	http.SetCookie(context.Writer, &http.Cookie{
		Name:     tokenString,
		Value:    token,
		MaxAge:   600,
		Path:     "/",
		Domain:   "localhost",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})
	context.Redirect(http.StatusMovedPermanently, MblogURI+HomeURI)
}

func GetSigninForm(context *gin.Context) {
	context.HTML(http.StatusOK, "signin.html", nil)
}

func GetSignupForm(context *gin.Context) {
	context.HTML(http.StatusOK, "signup.html", nil)
}
