package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	youtube    = "youtube"
	youtube2   = "youtu.be"
	youtubeURL = "https://www.youtube.com/embed/%s"

	pattern = "https:\\/\\/www.youtube.com\\/embed/\\.*"
)

type Obj map[string]interface{}

type Post struct {
	ID              uint64    `json:"id" binding:"required" bson:"_id"`
	Date            time.Time `json:"date" bson:"date"`
	Text            string    `json:"textContent" binding:"required"`
	ImageURL        string    `json:"imageURL" bson:"imageURL"`
	VideoURL        string    `json:"videoURL" bson:"videoURL"`
	YoutubeVideoURL string    `json:"youtubeVideoURL" bson:"youtubeVideoURL"`
	UserID          uint64    `json:"-" bson:"userID"`
}

func (post *Post) GetYoutubeVideoURL() error {
	if strings.Contains(post.VideoURL, youtube) || strings.Contains(post.VideoURL, youtube2) {
		URLparts := strings.Split(post.VideoURL, "/")
		videoURI := ""
		for _, part := range URLparts {
			if len(part) > len(videoURI) {
				switch {
				case !strings.Contains(part, youtube2):
					videoURI = strings.Split(post.VideoURL, "=")[1]
				case !strings.Contains(part, youtube):
					videoURI = part
				}
			}
		}
		result, err := CheckYoutubeVideoURL(fmt.Sprintf(youtubeURL, videoURI))
		if err != nil {
			return err
		}
		if result {
			post.YoutubeVideoURL = fmt.Sprintf(youtubeURL, videoURI)
			post.VideoURL = ""
		}
	}
	return nil
}

func CheckYoutubeVideoURL(url string) (bool, error) {
	return regexp.MatchString(pattern, url)
}

// redundant structs used for reading data from DB and showing it on web page
type PostWithEmail struct {
	ID              uint64    `json:"id" binding:"required" bson:"_id"`
	Date            time.Time `json:"date" bson:"date"`
	Text            string    `json:"textContent" bson:"text" binding:"required"`
	ImageURL        string    `json:"imageURL" bson:"imageURL"`
	VideoURL        string    `json:"videoURL" bson:"videoURL"`
	Email           string    `json:"email" bson:"email" binding:"required"`
	YoutubeVideoURL string    `json:"youtubeVideoURL" bson:"youtubeVideoURL"`
	UserID          uint64    `json:"-" bson:"userID" binding:"required"`
}

type PostWithEmbeddedUser struct {
	ID              uint64    `json:"id" binding:"required" bson:"_id"`
	Date            time.Time `json:"date" bson:"date"`
	Text            string    `json:"textContent" bson:"text" binding:"required"`
	ImageURL        string    `json:"imageURL" bson:"imageURL"`
	VideoURL        string    `json:"videoURL" bson:"videoURL"`
	YoutubeVideoURL string    `json:"youtubeVideoURL" bson:"youtubeVideoURL"`
	UserOwner       []Obj     `json:"userOwner" bson:"userOwner" binding:"required"`
	UserID          uint64    `json:"-" bson:"userID" binding:"required"`
}

func (post PostWithEmbeddedUser) CreatePostWithEmail() PostWithEmail {
	postEmail := PostWithEmail{
		Date:            post.Date,
		Text:            post.Text,
		ImageURL:        post.ImageURL,
		VideoURL:        post.VideoURL,
		YoutubeVideoURL: post.YoutubeVideoURL,
		Email:           post.UserOwner[0]["email"].(string),
	}
	return postEmail
}
