package usersdto

type CreateUserRequest struct {
	FullName  string `json:"fullName" form:"fullName" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	Role      string `json:"role"`
}

type UpdateUserRequest struct {
	FullName  string `json:"fullName" form:"fullName"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Role      string `json:"role"`
}