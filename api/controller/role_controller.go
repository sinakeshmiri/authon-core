package controller

import (
	"context"
	"errors"

	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
)

func (h *Handler) CreateRole(
	ctx context.Context,
	req api.CreateRoleRequestObject,
) (api.CreateRoleResponseObject, error) {
	ucReq := domain.CreateRoleRequest{
		Owner:       req.Body.Owner,
		Rollname:    req.Body.Rolename,
		Description: req.Body.Description,
	}
	err := h.roleUsecase.Create(ctx, &ucReq)
	if err == nil {
		return api.CreateRole201Response{}, nil
	}
	switch {
	case errors.Is(err, domain.ErrRoleAlreadyExists):
		return api.CreateRole400Response{}, nil
	default:
		return api.CreateRole500Response{}, err
	}
}
