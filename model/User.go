package model

import "time"

type User struct {
	ID string `gorm:"type:uuid;primaryKey"`

	Email string `gorm:"type:varchar(320);uniqueIndex;not null"`

	PasswordHash string `gorm:"type:text;not null;column:password_hash"`

	IsActive bool `gorm:"not null;default:true;column:is_active"`

	CreatedAt time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt time.Time `gorm:"not null;default:now();column:updated_at"`
}

func (User) TableName() string { return "users" }
