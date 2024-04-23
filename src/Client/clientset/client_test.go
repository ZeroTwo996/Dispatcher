package clientset

import (
	"context"
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewForConfig(t *testing.T) {
	clientset := DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}

	podList, err := clientset.CoreV1().Pods("cloudgame").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		t.Errorf("list pods error:%v\n", err)
	}

	fmt.Println("pod count:", len(podList.Items))
	for _, pod := range podList.Items {
		fmt.Printf("name: %s\n", pod.Name)
	}

}
