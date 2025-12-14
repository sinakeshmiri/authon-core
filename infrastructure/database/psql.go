package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenPostgres(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, fmt.Errorf("database dsn can not be empty")
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
