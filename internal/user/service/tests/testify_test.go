package tests

import (
	"MockTest/internal/user/dao"
	"MockTest/internal/user/dao/mocks"
	"MockTest/internal/user/service"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestifyEqual(t *testing.T) {
	//mock 准备工作
	mockUserDao := mocks.NewUserDao(t)
	mockUserDao.On("GetUserByMobile", context.Background(), "18").Return(&dao.User{
		Nickname: "bobby18",
	}, nil)

	//实际调用过程
	userServer := service.UserService{
		UserDao: mockUserDao,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	assert.Nil(t, err)
	assert.Equal(t, user.Nickname, "bobby18", "they should be equal")
}

func TestTestifyNotEqual(t *testing.T) {
	//mock 准备工作
	mockUserDao := mocks.NewUserDao(t)
	mockUserDao.On("GetUserByMobile", context.Background(), "18").Return(&dao.User{
		Nickname: "bobby19",
	}, nil)

	//实际调用过程
	userServer := service.UserService{
		UserDao: mockUserDao,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	assert.Nil(t, err)
	assert.NotEqual(t, user.Nickname, "bobby18", "they should be not equal")
}
