package clientset

import (
	"log"
	"path/filepath"
	"runtime"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func NewForConfig(filePath string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", filePath)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("连接失败")
	}

	return clientset

}

func DefaultClient() *kubernetes.Clientset {
	currentPath := getProjectRoot()

	configPath := filepath.Join(currentPath, "..", "..", "..", "..", "conf", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("连接失败")
	}

	return clientset

}

func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}
