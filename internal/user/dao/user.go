package dao

import (
	"context"
)

func (u *UserDaoImpl) GetUserByMobile(ctx context.Context, mobile string) (*User, error) {
	var user *User
	_ = u.db.Where(&User{Mobile: mobile}).First(&user)
	return user, nil
}

type User struct {
	Mobile   string
	Password string
	Nickname string
}
