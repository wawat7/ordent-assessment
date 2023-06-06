package response

import (
	"ordent-assessment/entity"
	"time"
)

type userResponse struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUser(user entity.User) *userResponse {
	return &userResponse{
		Id:        user.Id.Hex(),
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FormatUsers(users []entity.User) []*userResponse {
	formats := []*userResponse{}

	for _, user := range users {
		format := FormatUser(user)
		formats = append(formats, format)
	}
	return formats
}
