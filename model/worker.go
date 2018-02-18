package model

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/topfreegames/resources-check/extension"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

// Worker loops every interval and checks limits and
// requests from every Kubernetes controller
type Worker struct {
	config           *viper.Viper
	kubernetesClient kubernetes.Interface
	logger           logrus.FieldLogger
	interval         time.Duration
}

// NewWorker connects with Kubernetes API
// and returns a worker instance
func NewWorker(
	config *viper.Viper,
	kubernetesClientOrNil kubernetes.Interface,
	logger logrus.FieldLogger,
	inCluster bool,
	kubeconfigPath string,
) (*Worker, error) {
	worker := &Worker{
		config: config,
		logger: logger,
	}

	worker.loadConfigurationDefaults()
	worker.configureWorker()
	err := worker.configureKubernetesClient(
		kubernetesClientOrNil, inCluster, kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return worker, nil
}

func (w *Worker) loadConfigurationDefaults() {
	w.config.SetDefault("app.worker.interval", "1h")
}

func (w *Worker) configureWorker() {
	w.interval = w.config.GetDuration("app.worker.interval")
}

func (w *Worker) configureKubernetesClient(
	kubernetesClientOrNil kubernetes.Interface,
	inCluster bool,
	kubeconfigPath string,
) error {
	if kubernetesClientOrNil != nil {
		w.kubernetesClient = kubernetesClientOrNil
		return nil
	}
	clientset, err := extension.GetKubernetesClient(w.logger, inCluster, kubeconfigPath)
	if err != nil {
		return err
	}
	w.kubernetesClient = clientset
	return nil
}

// Start starts check every interval
func (w *Worker) Start() {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.logger.Infof("[%s] Checking controllers", currentTime())
			namespaces, err := w.listNamespaces()
			if err != nil {
				w.logger.WithError(err).Error("error listing namespaces")
			}

			fmt.Printf("%#v", namespaces)
		}
	}
}

func (w *Worker) listNamespaces() ([]string, error) {
	listOptions := v1.ListOptions{}

	kubeNamespaces, err := w.kubernetesClient.CoreV1().Namespaces().List(listOptions)
	if err != nil {
		return nil, err
	}
	namespaces := make([]string, len(kubeNamespaces.Items))

	for i, kubeNamespace := range kubeNamespaces.Items {
		namespaces[i] = kubeNamespace.Name
	}

	return namespaces, nil
}
