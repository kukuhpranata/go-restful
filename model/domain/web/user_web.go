package web

type CreateUserRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required" json:"id"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
}

type LoginUserRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type LoginUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
