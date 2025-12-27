package domain

import "errors"

var (
	ErrDatabaseQueryFailed        = errors.New("failed to execute the query")
	ErrUserAlreadyExists          = errors.New("user already exists")
	ErrPasswordHashCreationFailed = errors.New("failed to calculate password hash")
	ErrInvalidCredentials         = errors.New("invalid credentials")
	ErrRoleAlreadyExists          = errors.New("role already exists")
	ErrApplicationNotFound        = errors.New("application not found")
	ErrInvalidTransition          = errors.New("invalid application status transition")
)
