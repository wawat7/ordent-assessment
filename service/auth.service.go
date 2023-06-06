package service

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"ordent-assessment/config"
	"ordent-assessment/request/auth_request"
	"ordent-assessment/request/user_request"
)

type AuthService interface {
	Register(ctx context.Context, payload auth_request.RegisterRequest) error
	Login(ctx context.Context, payload auth_request.LoginRequest, configuration config.Config) (token string, err error)
	Logout(ctx context.Context, token string) error
	GenerateToken(userID string, configuration config.Config) (string, error)
	ValidateToken(token string, configuration config.Config) (*jwt.Token, error)
}

type authService struct {
	userService      UserService
	userTokenService UserTokenService
}

func NewAuthService(userService UserService, tokenService UserTokenService) *authService {
	return &authService{userService: userService, userTokenService: tokenService}
}

func (service *authService) Register(ctx context.Context, payload auth_request.RegisterRequest) error {

	user, _ := service.userService.GetByUsername(ctx, payload.Username)
	if user.Username != "" {
		return errors.New("username already taken")
	}
	userPermissions := []string{
		"can-read-product",
	}
	userPayload := user_request.CreateUserRequest{
		Username:    payload.Username,
		Password:    payload.Password,
		Name:        payload.Name,
		Permissions: userPermissions,
	}

	_, err := service.userService.Create(ctx, userPayload)
	if err != nil {
		return err
	}
	return nil

}

func (service *authService) Login(ctx context.Context, payload auth_request.LoginRequest, configuration config.Config) (token string, err error) {
	user, err := service.userService.GetByUsername(ctx, payload.Username)
	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return token, err
	}

	token, err = service.GenerateToken(user.Id.Hex(), configuration)
	if err != nil {
		return token, err
	}

	userToken, _ := service.userTokenService.FindToken(ctx, user.Id.Hex(), token)
	if userToken.Token == "" {
		err = service.userTokenService.Create(ctx, user.Id.Hex(), token)
		if err != nil {
			return token, err
		}
	}

	return token, nil
}

func (service *authService) Logout(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

func (service *authService) GenerateToken(userID string, configuration config.Config) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(configuration.Get("SECRET_KEY")))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (service *authService) ValidateToken(token string, configuration config.Config) (*jwt.Token, error) {
	token2, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(configuration.Get("SECRET_KEY")), nil
	})

	if err != nil {
		return token2, err
	}

	return token2, nil
}
