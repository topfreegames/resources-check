package model

import (
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
	monitors         []MonitorService
}

// NewWorker connects with Kubernetes API
// and returns a worker instance
func NewWorker(
	config *viper.Viper,
	kubernetesClientOrNil kubernetes.Interface,
	monitorsOrNil []MonitorService,
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
	err = worker.configureMonitors(monitorsOrNil)
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

func (w *Worker) configureMonitors(
	monitorsOrNil []MonitorService,
) error {
	if monitorsOrNil != nil {
		w.monitors = monitorsOrNil
		return nil
	}

	ddMonitor, err := NewDatadogMonitor(w.config)
	if err != nil {
		return err
	}
	w.monitors = []MonitorService{ddMonitor}
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

			for _, namespace := range namespaces {
				failedControllers, err := w.checkNamespace(namespace)
				if err != nil {
					w.logger.WithError(err).Error("error listing namespaces")
				}
				w.sendToMonitors(failedControllers)
			}
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

func (w *Worker) checkNamespace(
	namespace string,
) ([]string, error) {
	failedDeployments, err := w.checkDeployment(namespace)
	if err != nil {
		return nil, err
	}

	failedStatefulsets, err := w.checkStatefulset(namespace)
	if err != nil {
		return nil, err
	}

	failedDaemonsets, err := w.checkDaemonset(namespace)
	if err != nil {
		return nil, err
	}

	failedControllers := failedDeployments
	failedControllers = append(failedControllers, failedStatefulsets...)
	failedControllers = append(failedControllers, failedDaemonsets...)

	return failedControllers, nil
}

func (w *Worker) checkDeployment(
	namespace string,
) ([]string, error) {
	listOptions := v1.ListOptions{}
	deployments, err := w.kubernetesClient.
		Extensions().
		Deployments(namespace).
		List(listOptions)
	if err != nil {
		return nil, err
	}
	failedDeployments := []string{}
	for _, deployment := range deployments.Items {
		if !isIgnored(Deployment, &deployment, w.config) && !hasResources(deployment.Spec.Template.Spec.Containers) {
			failedDeployments = append(failedDeployments, name(&deployment))
		}
	}
	return failedDeployments, nil
}

func (w *Worker) checkStatefulset(
	namespace string,
) ([]string, error) {
	listOptions := v1.ListOptions{}
	statefulsets, err := w.kubernetesClient.
		Apps().
		StatefulSets(namespace).
		List(listOptions)
	if err != nil {
		return nil, err
	}
	failedStatefulsets := []string{}
	for _, statefulset := range statefulsets.Items {
		if !isIgnored(Statefulset, &statefulset, w.config) && !hasResources(statefulset.Spec.Template.Spec.Containers) {
			failedStatefulsets = append(failedStatefulsets, name(&statefulset))
		}
	}
	return failedStatefulsets, nil
}

func (w *Worker) checkDaemonset(
	namespace string,
) ([]string, error) {
	listOptions := v1.ListOptions{}
	daemonsets, err := w.kubernetesClient.
		Extensions().
		DaemonSets(namespace).
		List(listOptions)
	if err != nil {
		return nil, err
	}
	failedDaemonsets := []string{}
	for _, daemonset := range daemonsets.Items {
		if !isIgnored(Daemonset, &daemonset, w.config) && !hasResources(daemonset.Spec.Template.Spec.Containers) {
			failedDaemonsets = append(failedDaemonsets, name(&daemonset))
		}
	}
	return failedDaemonsets, nil
}

func (w *Worker) sendToMonitors(controllers []string) {
	for _, monitor := range w.monitors {
		err := monitor.Send(controllers...)
		if err != nil {
			w.logger.WithError(err).
				Errorf("failed to send to monitor service: %s", monitor.Name())
		}
	}
}
