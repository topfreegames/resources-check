package extension

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetKubernetesClient returns a Kubernetes client
func GetKubernetesClient(logger logrus.FieldLogger, inCluster bool, kubeConfigPath string) (kubernetes.Interface, error) {
	var err error
	l := logger.WithFields(logrus.Fields{
		"operation": "extension.GetKubernetesClient",
	})
	var config *rest.Config
	if inCluster {
		l.Debug("starting with incluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		l.Debug("starting outside Kubernetes cluster")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	}
	if err != nil {
		l.WithError(err).Error("start Kubernetes failed")
		return nil, err
	}
	l.Debug("connecting to Kubernetes...")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		l.WithError(err).Error("connection to Kubernetes failed")
		return nil, err
	}
	l.Info("successfully connected to Kubernetes")
	return clientset, nil
}
