package service

import (
	"MockTest/internal/user/dao"
	"context"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetUserByMobile(t *testing.T) {
	//mock 准备工作
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserData := dao.NewMockUserDao(ctrl)
	mockUserData.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(&dao.User{
		Nickname: "bobby18",
	}, nil)

	//实际调用过程
	userServer := UserService{
		userDao: mockUserData,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "18")

	//判断正确与否
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if user.Nickname != "bobby17" {
		t.Errorf("error: %v", err)
	}

}

func TestGetUserByMobileFail(t *testing.T) {
	//mock 准备工作
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserData := dao.NewMockUserDao(ctrl)
	mockUserData.EXPECT().GetUserByMobile(gomock.Any(), "19").Return(&dao.User{
		Nickname: "bobby19",
	}, nil)

	//实际调用过程
	userServer := UserService{
		userDao: mockUserData,
	}
	user, err := userServer.GetUserByMobile(context.Background(), "19")

	//判断正确与否
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if user.Nickname != "bobby19" {
		t.Errorf("error: %v", err)
	}

}
