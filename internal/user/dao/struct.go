package dao

import (
	"context"
	"gorm.io/gorm"
)

//go:generate mockgen -source=struct.go -destination=mock_struct.go -package=dao
type UserDao interface {
	GetUserByMobile(ctx context.Context, mobile string) (*User, error)
}

type UserDaoImpl struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDaoImpl {
	return &UserDaoImpl{db: db}
}
