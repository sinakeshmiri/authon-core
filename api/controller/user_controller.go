package controller

import (
	"context"
	"time"

	"github.com/google/uuid"
	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
	"github.com/sinakeshmiri/imcore/model"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	uc domain.CreateUserUsecase
}

func NewHandler(uc domain.CreateUserUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
func (h *Handler) CreateUser(
	ctx context.Context,
	req api.CreateUserRequestObject,
) (api.CreateUserResponseObject, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Body.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return api.CreateUser500Response{}, err
	}

	passwordHash := string(hash)
	id, err := uuid.NewUUID()
	if err != nil {
		return api.CreateUser500Response{}, err
	}
	newUser := model.User{
		ID:           id.String(),
		Email:        string(req.Body.Email),
		PasswordHash: passwordHash,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = h.uc.Create(ctx, &newUser)
	if err != nil {
		return api.CreateUser500Response{}, err
	}
	return api.CreateUser201Response{}, nil
}
