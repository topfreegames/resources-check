package testing

import (
	"github.com/topfreegames/resources-check/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	appsv1beta1 "k8s.io/client-go/pkg/apis/apps/v1beta1"
	extensionsv1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

// GetControllerFunc returns a controller type
type GetControllerFunc func(
	clientset kubernetes.Interface,
	name, namespace string,
	resources apiv1.ResourceRequirements,
) (controller.Controller, error)

func int32Ptr(i int32) *int32 { return &i }

//CreateNamespace creates a namespace
func CreateNamespace(clientset kubernetes.Interface, namespace string) error {
	ns := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
	_, err := clientset.CoreV1().Namespaces().Create(ns)
	return err
}

// CreateDeployment creates a deployment
func CreateDeployment(
	clientset kubernetes.Interface,
	name, namespace string,
	resources apiv1.ResourceRequirements,
) (controller.Controller, error) {
	deploymentsClient := clientset.AppsV1beta1().Deployments(namespace)
	deployment := &appsv1beta1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1beta1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Resources: resources,
						},
					},
				},
			},
		},
	}

	result, err := deploymentsClient.Create(deployment)
	return result, err
}

// CreateStatefulset creates a statefulset
func CreateStatefulset(
	clientset kubernetes.Interface,
	name, namespace string,
	resources apiv1.ResourceRequirements,
) (controller.Controller, error) {
	statefulsetsClient := clientset.AppsV1beta1().StatefulSets(namespace)
	statefulset := &appsv1beta1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1beta1.StatefulSetSpec{
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Resources: resources,
						},
					},
				},
			},
		},
	}

	result, err := statefulsetsClient.Create(statefulset)
	return result, err
}

// CreateDaemonset creates a statefulset
func CreateDaemonset(
	clientset kubernetes.Interface,
	name, namespace string,
	resources apiv1.ResourceRequirements,
) (controller.Controller, error) {
	daemonsetClient := clientset.ExtensionsV1beta1().DaemonSets(namespace)
	daemonset := &extensionsv1beta1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: extensionsv1beta1.DaemonSetSpec{
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Resources: resources,
						},
					},
				},
			},
		},
	}

	result, err := daemonsetClient.Create(daemonset)
	return result, err
}
