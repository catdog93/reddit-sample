// Service contains some logic. Used by controller.
package service

import (
	"errors"
	ai "github.com/night-codes/mgo-ai"
	mgo "gopkg.in/mgo.v2"
	"testTaskBitmediaLabs/entity"
)

const (
	DBURL               = "mongodb://localhost:2717"
	DBName              = "Mblog"
	UsersCollectionName = "Users"
)

var UsersCollection *mgo.Collection

type Obj map[string]interface{}

func CreateUser(user entity.User) error {
	user.ID = ai.Next(UsersCollectionName)
	user.UsersIDs = []uint64{user.ID}
	return UsersCollection.Insert(interface{}(user))
}

func FindUserByEmail(email string) (*entity.User, error) {
	result := entity.User{}
	err := UsersCollection.Find(Obj{"email": email}).One(&result)
	return &result, err
}

func Subscribe(subscriberUser *entity.User, emailOfPublisher string) (bool, error) {
	userSubscription, err := FindUserByEmail(emailOfPublisher)
	if err != nil {
		return false, err
	}
	isSubscribed, _ := CheckSubscription(subscriberUser.UsersIDs, userSubscription.ID)
	if isSubscribed {
		return true, nil
	}
	if subscriberUser.UsersIDs == nil {
		subscriberUser.UsersIDs = []uint64{userSubscription.ID}
	} else {
		subscriberUser.UsersIDs = append(subscriberUser.UsersIDs, userSubscription.ID)
	}

	user := Obj{
		"email":    subscriberUser.Email,
		"password": subscriberUser.Password,
		"usersIDs": subscriberUser.UsersIDs,
	}
	err = UsersCollection.UpdateId(subscriberUser.ID, user)
	return false, err
}

// also returns index
func CheckSubscription(userIDs []uint64, userID uint64) (bool, int) {
	if userIDs != nil {
		for index, id := range userIDs {
			if id == userID {
				return true, index
			}
		}
	}
	return false, -1
}

func Unfollow(subscriberUser *entity.User, emailOfPublisher string) (bool, error) {
	if subscriberUser.Email != emailOfPublisher {
		userSubscription, err := FindUserByEmail(emailOfPublisher)
		if err != nil {
			return false, err
		}
		isSubscribed, index := CheckSubscription(subscriberUser.UsersIDs, userSubscription.ID)
		if !isSubscribed {
			return true, nil
		}

		// quick removing one following from slice
		subscriberUser.UsersIDs[index] = subscriberUser.UsersIDs[len(subscriberUser.UsersIDs)-1]
		//subscriberUser.UsersIDs[len(subscriberUser.UsersIDs)-1] = 0
		subscriberUser.UsersIDs = subscriberUser.UsersIDs[:len(subscriberUser.UsersIDs)-1]

		user := Obj{
			"email":    subscriberUser.Email,
			"password": subscriberUser.Password,
			"usersIDs": subscriberUser.UsersIDs,
		}
		err = UsersCollection.UpdateId(subscriberUser.ID, user)
		return false, err
	}
	return false, errors.New("can't unsubscribed yourself")
}
