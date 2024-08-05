package dao

import (
	"context"
	"gorm.io/gorm"
)

//go:generate mockgen -source=struct.go -destination=mocks/gomock_struct.go -package=mocks
//go:generate mockery --name UserDao --filename testify_struct.go --outpkg mocks
type UserDao interface {
	GetUserByMobile(ctx context.Context, mobile string) (*User, error)
}

type UserDaoImpl struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDaoImpl {
	return &UserDaoImpl{db: db}
}
