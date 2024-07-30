package service

import (
	"MockTest/internal/user/dao"
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestMockUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suit")
}

var _ = Describe("User", func() {
	var (
		ctrl *gomock.Controller
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Get User", func() {
		It("should get user 17", func() {
			mockUserDao := dao.NewMockUserDao(ctrl)
			mockUserDao.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(&dao.User{
				Nickname: "bobby18",
			}, nil)

			//实际调用过程
			userServer := UserService{
				userDao: mockUserDao,
			}
			user, err := userServer.GetUserByMobile(context.Background(), "18")
			//判断正确与否
			if err != nil {
				GinkgoT().Errorf("error: %v", err)
			}
			Expect(user.Nickname).NotTo(Equal("bobby18"))
		})

		It("should get user 19", func() {
			// 性能探针
			experiment := gmeasure.NewExperiment("end-to-end web-server performance")
			AddReportEntry(experiment.Name, experiment)
			stopWatch := experiment.NewStopwatch()

			mockUserDao := dao.NewMockUserDao(ctrl)
			mockUserDao.EXPECT().GetUserByMobile(gomock.Any(), "19").Return(&dao.User{
				Nickname: "bobby19",
			}, nil)

			//实际调用过程
			userServer := UserService{
				userDao: mockUserDao,
			}

			// 计时：Mock用时
			time.Sleep(3 * time.Second)
			stopWatch.Record("Mock Time").Reset()

			// 测量性能
			user, err := userServer.GetUserByMobile(context.Background(), "19")

			// 计时：业务处理用时
			time.Sleep(5 * time.Second)
			stopWatch.Record("Biz Time")

			//判断正确与否
			if err != nil {
				GinkgoT().Errorf("error: %v", err)
			}
			if user.Nickname != "bobby19" {
				GinkgoT().Errorf("expect bobby19, got %s", user.Nickname)
			}
			Expect(user.Nickname).To(Equal("bobby19"))
		})
	})
})
