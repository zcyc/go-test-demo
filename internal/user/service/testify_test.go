package service_test

import (
	"MockTest/internal/user/dao"
	"MockTest/internal/user/service"
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTestifyEqual(t *testing.T) {
	//mock 准备工作
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserData := dao.NewMockUserDao(ctrl)
	mockUserData.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(&dao.User{
		Nickname: "bobby18",
	}, nil)

	//实际调用过程
	userServer := service.UserService{
		UserDao: mockUserData,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	assert.Nil(t, err)
	assert.Equal(t, user.Nickname, "bobby18", "they should be equal")
}

func TestTestifyNotEqual(t *testing.T) {
	//mock 准备工作
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserData := dao.NewMockUserDao(ctrl)
	mockUserData.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(&dao.User{
		Nickname: "bobby19",
	}, nil)

	//实际调用过程
	userServer := service.UserService{
		UserDao: mockUserData,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	assert.Nil(t, err)
	assert.NotEqual(t, user.Nickname, "bobby18", "they should be not equal")
}
