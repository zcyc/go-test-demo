package tests

import (
	"MockTest/pkg/user/dao"
	"MockTest/pkg/user/mocks"
	"MockTest/pkg/user/service"
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// 示例数据
var testUser = &dao.User{
	Mobile:   "13800138000",
	Nickname: "测试用户",
	Password: "password123",
}

// 标准 testing 包测试示例
func TestUserService(t *testing.T) {
	// 准备 mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserDao := mocks.NewMockUserDao(ctrl)
	mockUserDao.EXPECT().
		GetUserByMobile(gomock.Any(), "13800138000").
		Return(testUser, nil)

	// 创建服务实例
	userService := &service.UserService{
		UserDao: mockUserDao,
	}

	// 执行测试
	user, err := userService.GetUserByMobile(context.Background(), "13800138000")

	// 断言结果
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if user.Nickname != "测试用户" {
		t.Errorf("expected nickname '测试用户', got '%s'", user.Nickname)
	}
}

// Testify 测试示例
func TestUserServiceWithTestify(t *testing.T) {
	// 准备 mock
	mockUserDao := new(mocks.UserDao)
	mockUserDao.On("GetUserByMobile", context.Background(), "13800138000").
		Return(testUser, nil)

	// 创建服务实例
	userService := &service.UserService{
		UserDao: mockUserDao,
	}

	// 执行测试
	user, err := userService.GetUserByMobile(context.Background(), "13800138000")

	// 断言结果
	assert.NoError(t, err)
	assert.Equal(t, "测试用户", user.Nickname)
	assert.Equal(t, "13800138000", user.Mobile)
	mockUserDao.AssertExpectations(t)
}

// Ginkgo 测试套件
func TestUserServiceWithGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Service Suite")
}

// Ginkgo 测试规格
var _ = Describe("UserService", func() {
	var (
		ctrl        *gomock.Controller
		mockUserDao *mocks.MockUserDao
		userService *service.UserService
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockUserDao = mocks.NewMockUserDao(ctrl)
		userService = &service.UserService{
			UserDao: mockUserDao,
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("获取用户", func() {
		It("应该返回正确的用户信息", func() {
			// 设置 mock 预期
			mockUserDao.EXPECT().
				GetUserByMobile(gomock.Any(), "13800138000").
				Return(testUser, nil)

			// 执行测试
			user, err := userService.GetUserByMobile(context.Background(), "13800138000")

			// 断言结果
			Expect(err).To(BeNil())
			Expect(user.Nickname).To(Equal("测试用户"))
			Expect(user.Mobile).To(Equal("13800138000"))
		})
	})
})
