package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTaskBitmediaLabs/entity"
	"testTaskBitmediaLabs/service"
	"time"
)

const (
	CreatePostURI = "/postCreation"

	text        = "inputText"
	image       = "inputImage"
	video       = "inputVideo"
	tokenString = "token"
)

func CreatePost(context *gin.Context) {
	text := context.PostForm(text)
	imageURL := context.PostForm(image)
	videoURL := context.PostForm(video)

	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	userAuth, _ := service.TokensCache[token]
	user, err := service.FindUserByEmail(userAuth.Email)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	post := entity.Post{
		Text:     text,
		ImageURL: imageURL,
		VideoURL: videoURL,
		Date:     time.Now(),
		UserID:   user.ID,
	}
	err = post.GetYoutubeVideoURL()
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = service.CreatePost(post)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	}
	context.Redirect(http.StatusMovedPermanently, MblogURI+SubscriptionsURI)
}

func GetCreatePostForm(context *gin.Context) {
	token, err := context.Cookie(tokenString)
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
		return
	}
	user, _ := service.TokensCache[token]
	context.HTML(http.StatusOK, "createPostForm.html", user.Email)
}
