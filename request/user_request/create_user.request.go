package user_request

type CreateUserRequest struct {
	Username    string   `json:"username" binding:"requried"`
	Password    string   `json:"password" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Permissions []string `json:"permissions" binding:"required"`
}
