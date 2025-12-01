package user

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PaginationQuery struct {
	Page  int `form:"page,default=1"`
	Limit int `form:"Limit,default=1"`
}

type ListUsersResponse struct {
	Sucsses bool           `json:"success"`
	Message string         `json:"message"`
	Data    []UserResponse `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}
