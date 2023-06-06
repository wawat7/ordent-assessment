package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"ordent-assessment/entity"
	"ordent-assessment/repository"
	"ordent-assessment/request/user_request"
)

type UserService interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetById(ctx context.Context, Id string) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	Create(ctx context.Context, payload user_request.CreateUserRequest) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (service *userService) GetAll(ctx context.Context) ([]entity.User, error) {
	users, err := service.repo.FindAll(ctx)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (service *userService) GetById(ctx context.Context, Id string) (entity.User, error) {
	user, err := service.repo.FindById(ctx, Id)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (service *userService) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	user, err := service.repo.FindByUsername(ctx, username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *userService) Create(ctx context.Context, payload user_request.CreateUserRequest) (string, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	user := entity.User{
		Username:    payload.Username,
		Password:    string(passwordHash),
		Name:        payload.Name,
		Permissions: payload.Permissions,
	}
	userId, err := service.repo.Create(ctx, user)
	if err != nil {
		return userId, err
	}
	return userId, nil
}
