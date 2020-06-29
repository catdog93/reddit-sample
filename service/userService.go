// Service contains some logic. Used by controller.
package service

import (
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

func Subscribe(currentUser *entity.User, email string) error {
	userSubscription, err := FindUserByEmail(email)
	if err != nil {
		return err
	}
	isSubscribed, _ := CheckSubscription(currentUser.UsersIDs, userSubscription.ID)
	if isSubscribed {
		return nil
	}
	if currentUser.UsersIDs == nil {
		currentUser.UsersIDs = []uint64{userSubscription.ID}
	} else {
		currentUser.UsersIDs = append(currentUser.UsersIDs, userSubscription.ID)
	}

	user := Obj{
		"email":    currentUser.Email,
		"password": currentUser.Password,
		"usersIDs": currentUser.UsersIDs,
	}
	err = UsersCollection.UpdateId(currentUser.ID, user)
	return err
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

func Unfollow(currentUser *entity.User, email string) error {
	if currentUser.Email != email {
		userSubscription, err := FindUserByEmail(email)
		if err != nil {
			return err
		}
		isSubscribed, index := CheckSubscription(currentUser.UsersIDs, userSubscription.ID)
		if !isSubscribed {
			return nil
		}

		// quick removing one following from slice
		currentUser.UsersIDs[index] = currentUser.UsersIDs[len(currentUser.UsersIDs)-1]
		currentUser.UsersIDs[len(currentUser.UsersIDs)-1] = 0
		currentUser.UsersIDs = currentUser.UsersIDs[:len(currentUser.UsersIDs)-1]

		user := Obj{
			"email":    currentUser.Email,
			"password": currentUser.Password,
			"usersIDs": currentUser.UsersIDs,
		}
		err = UsersCollection.UpdateId(currentUser.ID, user)
		return err
	}
	return nil
}
