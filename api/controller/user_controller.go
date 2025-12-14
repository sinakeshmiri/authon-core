package controller

import (
	"context"
	"crypto/sha256"
	"time"

	"github.com/google/uuid"
	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
	"github.com/sinakeshmiri/imcore/model"
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
	hashed := sha256.New()
	hashed.Write([]byte(req.Body.Password))
	hashedPassword := hashed.Sum(nil)
	id, err := uuid.NewUUID()
	if err != nil {
		return api.CreateUser500Response{}, err
	}
	newUser := model.User{
		ID:           id.String(),
		Email:        string(req.Body.Email),
		PasswordHash: string(hashedPassword),
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
