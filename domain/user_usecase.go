package domain

import (
	"context"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Fullname string `json:"fullname" binding:"required,fullname"`
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required"`
}

type UserUsecase interface {
	Create(c context.Context, req *CreateUserRequest) error
}
