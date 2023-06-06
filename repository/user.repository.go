package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ordent-assessment/entity"
	"time"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	Create(ctx context.Context, user entity.User) (string, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *userRepository {
	return &userRepository{collection: database.Collection("users")}
}

func (repo *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	results, err := repo.collection.Find(ctx, bson.M{"deleted_at": nil})
	if err != nil {
		return users, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var user entity.User
		err := results.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepository) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := repo.collection.FindOne(ctx, bson.M{"username": username, "deleted_at": nil}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) Create(ctx context.Context, user entity.User) (string, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := repo.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
