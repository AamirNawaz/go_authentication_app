package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MUser struct {
	Id       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password []byte             `json:"password"`
	Status   string             `json:"status"`
}
