package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ordent-assessment/entity"
	"ordent-assessment/repository"
)

type UserTokenService interface {
	FindToken(ctx context.Context, userId string, token string) (entity.UserToken, error)
	Create(ctx context.Context, userId string, token string) error
	Delete(ctx context.Context, userId string, token string) error
}

type userTokenService struct {
	repo repository.UserTokenRepository
}

func NewUserTokenService(repo repository.UserTokenRepository) *userTokenService {
	return &userTokenService{repo: repo}
}

func (service *userTokenService) FindToken(ctx context.Context, userId string, token string) (entity.UserToken, error) {
	userToken, err := service.repo.FindToken(ctx, userId, token)
	if err != nil {
		return userToken, err
	}
	return userToken, nil
}

func (service *userTokenService) Create(ctx context.Context, userId string, token string) error {
	objId, _ := primitive.ObjectIDFromHex(userId)
	userToken := entity.UserToken{
		UserId: objId,
		Token:  token,
	}

	err := service.repo.Create(ctx, userToken)
	if err != nil {
		return err
	}
	return nil
}

func (service *userTokenService) Delete(ctx context.Context, userId string, token string) error {
	userToken, err := service.FindToken(ctx, userId, token)
	if err != nil {
		return err
	}

	err = service.repo.Delete(ctx, userToken)
	if err != nil {
		return err
	}
	return nil
}
