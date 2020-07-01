package service

import (
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"testTaskBitmediaLabs/entity"
)

const (
	PostsCollectionName = "Posts"
)

var PostsCollection *mgo.Collection

var QueryHome = []Obj{
	{
		"$lookup": Obj{
			"from":         UsersCollectionName,
			"localField":   "userID",
			"foreignField": "_id",
			"as":           "userOwner",
		},
	},
}

func CreatePost(post entity.Post) error {
	post.ID = ai.Next(PostsCollectionName)
	return PostsCollection.Insert(interface{}(post))
}

// if ids == nil, function returns all posts from DB. Otherwise returns posts of specified users
func GetPostsWithEmail(userIDs ...uint64) ([]entity.PostWithEmail, error) {
	var customPosts []entity.PostWithEmbeddedUser
	var pipe *mgo.Pipe
	pipe = PostsCollection.Pipe(QueryHome)
	err := pipe.All(&customPosts)
	if err != nil {
		return nil, err
	}
	postsWithEmail := make([]entity.PostWithEmail, 0, len(customPosts))
	if userIDs != nil {
		for _, value := range customPosts {
			for _, id := range userIDs {
				if value.UserID == id {
					postsWithEmail = append(postsWithEmail, value.CreatePostWithEmail())
				}
			}
		}
		return postsWithEmail, nil
	}
	for _, value := range customPosts {
		postsWithEmail = append(postsWithEmail, value.CreatePostWithEmail())
	}
	return postsWithEmail, nil
}
