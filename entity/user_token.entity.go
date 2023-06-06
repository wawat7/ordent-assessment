package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserToken struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	UserId primitive.ObjectID `bson:"user_id,omitempty"`
	Token  string             `bson:"token,omitempty"`
}
