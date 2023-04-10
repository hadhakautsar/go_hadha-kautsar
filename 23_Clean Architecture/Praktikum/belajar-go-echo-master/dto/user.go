package dto

type UserRequest struct {
	ID       int    `json:"id" form:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
