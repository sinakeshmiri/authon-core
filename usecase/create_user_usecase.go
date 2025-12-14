package usecase

import (
	"context"
	"time"

	"github.com/sinakeshmiri/imcore/domain"
	"github.com/sinakeshmiri/imcore/model"
)

type creatUserUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (cu creatUserUsecase) Create(ctx context.Context, user *model.User) error {
	err := cu.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (cu creatUserUsecase) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	byEmail, err := cu.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return *byEmail, nil
}

func NewCreateUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.CreateUserUsecase {
	return &creatUserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
