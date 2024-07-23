package mock

import (
	mock "MockTest/internal/user/service"
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
			mockUserData := NewMockUserData(ctrl)
			mockUserData.EXPECT().GetUserByMobile(gomock.Any(), "18").Return(mock.User{
				NickName: "bobby18",
			}, nil)

			//实际调用过程
			userServer := mock.UserServer{
				Db: mockUserData,
			}
			user, err := userServer.GetUserByMobile(context.Background(), "18")
			//判断正确与否
			if err != nil {
				GinkgoT().Errorf("error: %v", err)
			}
			Expect(user.NickName).NotTo(Equal("bobby18"))
		})

		It("should get user 19", func() {
			// 性能探针
			experiment := gmeasure.NewExperiment("end-to-end web-server performance")
			AddReportEntry(experiment.Name, experiment)
			stopWatch := experiment.NewStopwatch()

			mockUserData := NewMockUserData(ctrl)
			mockUserData.EXPECT().GetUserByMobile(gomock.Any(), "19").Return(mock.User{
				NickName: "bobby19",
			}, nil)

			//实际调用过程
			userServer := mock.UserServer{
				Db: mockUserData,
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
			if user.NickName != "bobby19" {
				GinkgoT().Errorf("expect bobby19, got %s", user.NickName)
			}
			Expect(user.NickName).To(Equal("bobby19"))
		})
	})

	DescribeTable("多测试用例", func(x, y int, expected bool) {
		Expect(x > y).To(Equal(expected))
	},
		Entry("x > y", 20, 10, true),
		Entry("x = y", 0, 0, false),
		Entry("x < y", 0, 10, false),
	)
})
