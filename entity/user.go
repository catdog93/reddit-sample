package entity

type User struct {
	ID       uint64   `json:"id" binding:"required" bson:"_id"`
	Email    string   `json:"email" bson:"email" binding:"required"`
	Password string   `json:"password" bson:"password" binding:"required"`
	UsersIDs []uint64 `json:"usersIDs" bson:"usersIDs"`
}

type UserBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Email struct {
	Email string `json:"email" bson:"email"`
}

type Message struct {
	Message string `json:"message"`
}
