package dao

import (
	"context"
)

//go:generate mockgen -source=struct.go -destination=../mocks/gomock_struct.go -package=mocks
//go:generate mockery --name UserDao --filename testify_struct.go --outpkg mocks --output ../mocks
type UserDao interface {
	GetUserByMobile(ctx context.Context, mobile string) (*User, error)
}

type UserDaoImpl struct {
	// 不再使用数据库
}

func NewUserDao() *UserDaoImpl {
	return &UserDaoImpl{}
}
