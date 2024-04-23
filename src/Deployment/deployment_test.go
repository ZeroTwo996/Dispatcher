package deployment

import (
	"testing"

	client "Dispatcher/src/client/clientset"

	v1 "k8s.io/api/core/v1"
)

func TestCreateDeployment(t *testing.T) {
	clientset := client.DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}

	affinity := &v1.Affinity{
		NodeAffinity: &v1.NodeAffinity{
			RequiredDuringSchedulingIgnoredDuringExecution: &v1.NodeSelector{
				NodeSelectorTerms: []v1.NodeSelectorTerm{
					{
						MatchExpressions: []v1.NodeSelectorRequirement{
							{
								Key:      "Zone_id",
								Operator: v1.NodeSelectorOpIn,
								Values:   []string{"HuaDong"},
							},
							{
								Key:      "Site_id",
								Operator: v1.NodeSelectorOpIn,
								Values:   []string{"NingBo"},
							},
							{
								Key:      "Server_ip",
								Operator: v1.NodeSelectorOpIn,
								Values:   []string{"10.10.103.50"},
							},
							// {
							// 	Key:      "role",
							// 	Operator: v1.NodeSelectorOpIn,
							// 	Values:   []string{"central"},
							// },
						},
					},
				},
			},
		},
	}
	err := CreateDeployment(clientset, "cloudgame", "cloudgame-ningbo-node01", "nginx", "nginx:latest", 5, affinity)
	if err != nil {
		t.Fatalf("Failed to create deployment: %v", err)
	}
}

func TestDeleteDeployment(t *testing.T) {
	clientset := client.DefaultClient()
	if clientset == nil {
		t.Error("Expected non-nil clientset, but got nil")
	}
	err := DeleteDeployment(clientset, "cloudgame", "nginx-deployment-ningbo01")
	if err != nil {
		t.Fatalf("Failed to create namespace: %v", err)
	}
}
