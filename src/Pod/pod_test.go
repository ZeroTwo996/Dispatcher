package pod

import (
	client "Dispatcher/src/client/clientset"
	"testing"
)

func TestCreatePod(t *testing.T) {
	clientset := client.DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}

	// 创建 PodInfo 实例
	podInfo := &PodInfo{
		PodName:       "test-pod",
		Namespace:     "cloudgame",
		ContainerName: "nginx",
		Image:         "nginx:latest",
		Labels:        nil,
		// NodeSelector:  nil,
		NodeSelector: map[string]string{"Zone_id": "HuaDong", "role": "central"},
	}
	err := CreatePod(podInfo)
	if err != nil {
		t.Fatalf("Failed to create pod: %v", err)
	}
}

func TestDeletePod(t *testing.T) {
	clientset := client.DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}

	err := DeletePod("cloudgame", "test-pod")
	if err != nil {
		t.Fatalf("Failed to delete pod: %v", err)
	}
}
