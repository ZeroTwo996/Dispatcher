package namespace

import (
	"testing"

	client "Dispatcher/src/client/clientset"
)

func TestCreateNamespace(t *testing.T) {
	clientset := client.DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}
	err := CreateNamespace(clientset, "test-namespace")
	if err != nil {
		t.Fatalf("Failed to create namespace: %v", err)
	}
}

func TestDeleteNamespace(t *testing.T) {
	clientset := client.DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}
	err := DeleteNamespace(clientset, "test-namespace")
	if err != nil {
		t.Fatalf("Failed to create namespace: %v", err)
	}
}
