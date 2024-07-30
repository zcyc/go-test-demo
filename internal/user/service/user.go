package service

import (
	"MockTest/internal/user/dao"
	"context"
)

/*
gomock 测试
安装：
	go get -u go.uber.org/mock

	go install go.uber.org/mock/mockgen@latest
*/

type UserService struct {
	userDao dao.UserDao
}

func (s *UserService) GetUserByMobile(ctx context.Context, mobile string) (*dao.User, error) {
	user, err := s.userDao.GetUserByMobile(ctx, mobile)
	if err != nil {
		return user, err
	}
	if user.Nickname == "bobby18" {
		user.Nickname = "bobby17"
	}
	return user, nil
}
