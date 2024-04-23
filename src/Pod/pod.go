package pod

import (
	client "Dispatcher/src/client/clientset"
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodInfo struct {
	PodName       string            `json:"podName"`
	Namespace     string            `json:"namespace"`
	ContainerName string            `json:"containerName"`
	Image         string            `json:"image"`
	Labels        map[string]string `json:"labels"`
	NodeSelector  map[string]string `json:"nodeSelector"`
}

func CreatePod(podInfo *PodInfo) error {
	client := client.DefaultClient()

	// 创建 Pod 对象
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podInfo.PodName,
			Namespace: podInfo.Namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  podInfo.ContainerName,
					Image: podInfo.Image,
				},
			},
			NodeSelector: podInfo.NodeSelector,
		},
	}
	_, err := client.CoreV1().Pods(podInfo.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	return err
}

func DeletePod(namespace string, podName string) error {
	client := client.DefaultClient()
	err := client.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	return err
}
