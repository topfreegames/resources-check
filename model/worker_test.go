package model_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topfreegames/resources-check/controller"
	. "github.com/topfreegames/resources-check/model"

	"github.com/topfreegames/resources-check/testing"
	"k8s.io/apimachinery/pkg/api/resource"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

var _ = Describe("Worker", func() {
	var worker *Worker
	var err error
	var name, namespace = "test", "test"
	var stopAfter = func() {
		time.Sleep(100 * time.Millisecond)
		worker.Run = false
	}

	var getControllerFuncs = []testing.GetControllerFunc{
		testing.CreateDeployment,
		testing.CreateStatefulset,
		testing.CreateDaemonset,
	}

	BeforeEach(func() {
		worker, err = NewWorker(config, clientset,
			[]MonitorService{mockMonitor}, logger, false, "")
		Expect(err).NotTo(HaveOccurred())

		err := testing.CreateNamespace(clientset, namespace)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Start", func() {
		for _, getController := range getControllerFuncs {
			It("should start", func() {
				go stopAfter()
				worker.Start()
			})

			It("should report error if controller has no requests and no limits", func() {
				kubeController, err := getController(clientset, name,
					namespace, apiv1.ResourceRequirements{})
				Expect(err).NotTo(HaveOccurred())

				mockMonitor.EXPECT().Send(controller.Name(kubeController))
				mockMonitor.EXPECT().Send(controller.Name(kubeController)).Do(func(...string) {
					worker.Run = false
				})

				worker.Start()
			})

			It("should report error if controller has only request", func() {
				cpu, err := resource.ParseQuantity("1m")
				Expect(err).NotTo(HaveOccurred())
				memory, err := resource.ParseQuantity("1Mi")
				Expect(err).NotTo(HaveOccurred())

				resource := apiv1.ResourceRequirements{
					Requests: apiv1.ResourceList{
						apiv1.ResourceRequestsCPU:    cpu,
						apiv1.ResourceRequestsMemory: memory,
					},
				}
				kubeController, err := getController(clientset, name,
					namespace, resource)
				Expect(err).NotTo(HaveOccurred())

				mockMonitor.EXPECT().Send(controller.Name(kubeController))
				mockMonitor.EXPECT().Send(controller.Name(kubeController)).Do(func(...string) {
					worker.Run = false
				})

				worker.Start()
			})

			It("should not report error if controller has only limits", func() {
				cpu, err := resource.ParseQuantity("1m")
				Expect(err).NotTo(HaveOccurred())
				memory, err := resource.ParseQuantity("1Mi")
				Expect(err).NotTo(HaveOccurred())

				resource := apiv1.ResourceRequirements{
					Limits: apiv1.ResourceList{
						apiv1.ResourceLimitsCPU:    cpu,
						apiv1.ResourceLimitsMemory: memory,
					},
				}
				_, err = getController(clientset, name,
					namespace, resource)
				Expect(err).NotTo(HaveOccurred())

				go stopAfter()
				worker.Start()
			})

			It("should not report error if controller has requests and limits", func() {
				cpu, err := resource.ParseQuantity("1m")
				Expect(err).NotTo(HaveOccurred())
				memory, err := resource.ParseQuantity("1Mi")
				Expect(err).NotTo(HaveOccurred())

				resource := apiv1.ResourceRequirements{
					Requests: apiv1.ResourceList{
						apiv1.ResourceRequestsCPU:    cpu,
						apiv1.ResourceRequestsMemory: memory,
					},
					Limits: apiv1.ResourceList{
						apiv1.ResourceLimitsCPU:    cpu,
						apiv1.ResourceLimitsMemory: memory,
					},
				}
				_, err = getController(clientset, name,
					namespace, resource)
				Expect(err).NotTo(HaveOccurred())

				go stopAfter()
				worker.Start()
			})
		}
	})
})
