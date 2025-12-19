package controller

import (
	"context"
	"errors"

	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
)

func (h *Handler) CreateUser(
	ctx context.Context,
	req api.CreateUserRequestObject,
) (api.CreateUserResponseObject, error) {
	ucReq := domain.CreateUserRequest{
		Fullname: req.Body.Fullname,
		Username: req.Body.Username,
		Email:    string(req.Body.Email),
		Password: req.Body.Password,
	}
	err := h.userUsecase.Create(ctx, &ucReq)
	if err == nil {
		return api.CreateUser201Response{}, nil
	}
	switch {
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return api.CreateUser400Response{}, nil
	default:
		return api.CreateUser500Response{}, err
	}
}
