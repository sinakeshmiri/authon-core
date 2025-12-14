package domain

import (
	"context"

	"github.com/sinakeshmiri/imcore/model"
)

type UserRepository interface {
	Create(c context.Context, user *model.User) error
	FindByEmail(c context.Context, email string) (*model.User, error)
}
