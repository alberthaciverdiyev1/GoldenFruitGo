package dto

type UserLoginRequest struct {
	UserName string `json:"user_name" binding:"required,min=2" form:"user_name"`
	Password string `json:"password" binding:"required,min=6" form:"password"`
}

type UserResponse struct {
	UserID   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}
