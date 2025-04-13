package dao

import (
	"context"
	"errors"
)

type User struct {
	Mobile   string
	Password string
	Nickname string
}

// 模拟内存数据
var users = map[string]*User{
	"13800138000": {
		Mobile:   "13800138000",
		Password: "password123",
		Nickname: "测试用户",
	},
}

func (u *UserDaoImpl) GetUserByMobile(ctx context.Context, mobile string) (*User, error) {
	if user, ok := users[mobile]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}
