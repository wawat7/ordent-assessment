package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Username    string             `bson:"username,omitempty"`
	Password    string             `bson:"password,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Permissions []string           `bson:"permissions,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `bson:"deleted_at,omitempty"`
}
