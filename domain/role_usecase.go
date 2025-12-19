package domain

import (
	"context"
)

type CreateRoleRequest struct {
	Rollname    string `json:"rolename" binding:"required,rolename"`
	Owner       string `json:"OwnerUsername" binding:"required,owner"`
	Description string `json:"description"`
}

type RoleUsecase interface {
	Create(c context.Context, req *CreateRoleRequest) error
}
