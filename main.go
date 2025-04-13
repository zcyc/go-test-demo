package main

import (
	"MockTest/pkg/user/dao"
	"MockTest/pkg/user/service"
	"context"
	"fmt"
)

// 模拟DAO实现
type mockUserDao struct{}

func (m *mockUserDao) GetUserByMobile(ctx context.Context, mobile string) (*dao.User, error) {
	// 模拟返回用户数据
	return &dao.User{
		Mobile:   "13800138000",
		Nickname: "测试用户",
		Password: "password123",
	}, nil
}

func main() {
	// 使用模拟的DAO实现
	mockDao := &mockUserDao{}

	// 创建用户服务，注入模拟DAO
	userService := &service.UserService{
		UserDao: mockDao,
	}

	// 查询用户
	user, err := userService.GetUserByMobile(context.Background(), "13800138000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印用户信息
	if user != nil && user.Mobile != "" {
		fmt.Printf("用户信息: 手机号=%s, 昵称=%s\n", user.Mobile, user.Nickname)
	} else {
		fmt.Println("用户不存在")
	}
}
