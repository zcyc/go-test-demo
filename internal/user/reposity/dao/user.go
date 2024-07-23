package mysql

import (
	mock "MockTest/internal/user/service"
	"context"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (u *User) GetUserByMobile(ctx context.Context, mobile string) (mock.User, error) {
	var user mock.User
	_ = u.db.Where(&mock.User{Mobile: mobile}).First(&user)
	return user, nil
}

var _ mock.UserData = &User{}
