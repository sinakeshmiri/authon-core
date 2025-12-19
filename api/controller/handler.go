package controller

import (
	"github.com/sinakeshmiri/imcore/domain"
)

func NewHandler(userUsecase domain.UserUsecase, roleUsecase domain.RoleUsecase) *Handler {
	return &Handler{
		userUsecase: userUsecase,
		roleUsecase: roleUsecase,
	}
}

type Handler struct {
	userUsecase domain.UserUsecase
	roleUsecase domain.RoleUsecase
}
