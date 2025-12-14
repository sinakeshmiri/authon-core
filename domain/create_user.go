package domain

import (
	"context"

	"github.com/sinakeshmiri/imcore/model"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserUsecase interface {
	Create(c context.Context, user *model.User) error
	GetUserByEmail(c context.Context, email string) (model.User, error)
}
