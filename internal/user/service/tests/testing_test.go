package tests

import (
	"MockTest/internal/user/dao"
	"MockTest/internal/user/dao/mocks"
	"MockTest/internal/user/service"
	"context"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTestingEqual(t *testing.T) {
	//mock 准备工作
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserDao := mocks.NewMockUserDao(ctrl)
	mockUserDao.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(&dao.User{
		Nickname: "bobby18",
	}, nil)

	//实际调用过程
	userServer := service.UserService{
		UserDao: mockUserDao,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if user.Nickname != "bobby18" {
		t.Errorf("error: %v", err)
	}

}

func TestTestingNotEqual(t *testing.T) {
	//mock 准备工作
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserDao := mocks.NewMockUserDao(ctrl)
	mockUserDao.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(&dao.User{
		Nickname: "bobby19",
	}, nil)

	//实际调用过程
	userServer := service.UserService{
		UserDao: mockUserDao,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if user.Nickname == "bobby18" {
		t.Errorf("error: %v", err)
	}

}
