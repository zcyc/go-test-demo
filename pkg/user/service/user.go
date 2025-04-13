package service

import (
	"MockTest/pkg/user/dao"
	"context"
)

type UserService struct {
	UserDao dao.UserDao
}

func (s *UserService) GetUserByMobile(ctx context.Context, mobile string) (*dao.User, error) {
	user, err := s.UserDao.GetUserByMobile(ctx, mobile)
	if err != nil {
		return user, err
	}
	return user, nil
}
