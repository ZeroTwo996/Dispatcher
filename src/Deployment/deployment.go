package deployment

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentInfo struct {
	Namespace      string `json:"namespace"`
	DeploymentName string `json:"deploymentName"`
	ContainerName  string `json:"containerName"`
	Image          string `json:"image"`
	Replicas       int32  `json:"replicas"`
}

// Helper function to convert int32 to pointer
func int32Ptr(i int32) *int32 {
	return &i
}

func CreateDeployment(clientset *kubernetes.Clientset, namespace, deploymentName, containerName, image string, replicas int32, affinity *v1.Affinity) error {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": deploymentName},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": deploymentName}},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  containerName,
							Image: image,
						},
					},
					Affinity: affinity,
				},
			},
		},
	}
	_, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	return err
}

func DeleteDeployment(clientset *kubernetes.Clientset, namespace string, deploymentName string) error {
	err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	return err
}
