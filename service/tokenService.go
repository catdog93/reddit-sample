package service

import (
	"github.com/dgrijalva/jwt-go"
	"testTaskBitmediaLabs/entity"
	"time"
)

const (
	SecretWord     = "I like my mblog"
	authorized     = "authorized"
	expired        = "exp"
	ExpirationTime = 15
)

var TokensCache = map[string]entity.User{}

func CreateToken(user *entity.User) string {
	jwtToken, err := generateToken()
	if err != nil {
		return err.Error()
	}
	TokensCache[jwtToken] = *user
	return jwtToken
}

func generateToken() (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims[authorized] = true
	atClaims[expired] = time.Now().Add(time.Minute * ExpirationTime).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(SecretWord))
	if err != nil {
		return "", err
	}
	return token, nil
}
