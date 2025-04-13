package tests

import (
	"MockTest/pkg/user/dao"
	"MockTest/pkg/user/mocks"
	"MockTest/pkg/user/service"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestifyEqual(t *testing.T) {
	//mock 准备工作
	mockUserDao := new(mocks.UserDao)
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
	mockUserDao := new(mocks.UserDao)
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
