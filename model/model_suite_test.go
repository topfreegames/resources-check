package model_test

import (
	"testing"

	"k8s.io/client-go/kubernetes/fake"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/topfreegames/resources-check/model"
	rtest "github.com/topfreegames/resources-check/testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var (
	config      *viper.Viper
	err         error
	clientset   *fake.Clientset
	hook        *test.Hook
	logger      *logrus.Logger
	mockMonitor *model.MockMonitorService
	mockCtrl    *gomock.Controller
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}

var _ = BeforeEach(func() {
	config, err = rtest.GetDefaultConfig()
	clientset = fake.NewSimpleClientset()
	logger, hook = test.NewNullLogger()
	logger.Level = logrus.DebugLevel

	mockCtrl = gomock.NewController(GinkgoT())
	mockMonitor = model.NewMockMonitorService(mockCtrl)
})
