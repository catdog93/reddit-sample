package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"testTaskBitmediaLabs/service"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func InitOAuth() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/users/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func GoogleLogin(context *gin.Context) {
	InitOAuth()
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	context.Redirect(http.StatusTemporaryRedirect, url)
}

func SignInViaGoogle(context *gin.Context) {
	token, err := service.GenerateToken()
	if err != nil {
		context.String(http.StatusInternalServerError, err.Error())
	}
	http.SetCookie(context.Writer, &http.Cookie{
		Name:     googleTokenString,
		Value:    token,
		MaxAge:   service.ExpirationTime * 60,
		Path:     "/",
		Domain:   "localhost",
		SameSite: http.SameSiteStrictMode,
		Secure:   false,
		HttpOnly: true,
	})
	context.Redirect(http.StatusMovedPermanently, MblogURI+HomeURI)

	//stateString, ok1 := context.GetPostForm("state")
	//codeString, ok2 := context.GetPostForm("code")
	//if ok1 && ok2 {
	//	content, err := getUserInfo(stateString, codeString)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		context.Redirect(http.StatusTemporaryRedirect, SigninURI)
	//		return
	//	}
	//
	//	context.String(http.StatusOK, "Content: %s\n", content)
	//}
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
