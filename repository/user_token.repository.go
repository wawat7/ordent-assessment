package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ordent-assessment/entity"
)

type UserTokenRepository interface {
	FindToken(ctx context.Context, userId string, token string) (entity.UserToken, error)
	Create(ctx context.Context, userToken entity.UserToken) error
	Delete(ctx context.Context, userToken entity.UserToken) error
}

type userTokenRepository struct {
	collection *mongo.Collection
}

func NewUserTokenRepository(database *mongo.Database) *userTokenRepository {
	return &userTokenRepository{collection: database.Collection("user_tokens")}
}

func (repo *userTokenRepository) FindToken(ctx context.Context, userId string, token string) (entity.UserToken, error) {
	var userToken entity.UserToken

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := repo.collection.FindOne(ctx, bson.M{"user_id": objId, "token": token}).Decode(&userToken)
	if err != nil {
		return userToken, err
	}
	return userToken, nil
}

func (repo *userTokenRepository) Create(ctx context.Context, userToken entity.UserToken) error {
	_, err := repo.collection.InsertOne(ctx, userToken)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userTokenRepository) Delete(ctx context.Context, userToken entity.UserToken) error {
	_, err := repo.collection.DeleteOne(ctx, userToken)
	if err != nil {
		return err
	}
	return nil
}
