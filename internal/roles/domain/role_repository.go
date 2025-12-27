package domain

import (
	"context"
	"time"
)

type RoleRepository interface {
	Create(c context.Context, role *Role) error
	FindByName(c context.Context, name string) (*Role, error)
}

type Role struct {
	Rolename      string    `gorm:"type:varchar(320);uniqueIndex;not null;primaryKey"`
	OwnerUsername string    `gorm:"type:varchar(320);not null"`
	Description   string    `gorm:"type:varchar(640);"`
	IsActive      bool      `gorm:"not null;default:true;column:is_active"`
	CreatedAt     time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt     time.Time `gorm:"not null;default:now();column:updated_at"`
}

func (Role) TableName() string { return "roles" }
