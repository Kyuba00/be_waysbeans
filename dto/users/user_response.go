package usersdto

type UserResponse struct {
	ID       int    `json:"ID" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
