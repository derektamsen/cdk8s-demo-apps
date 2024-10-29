package main

import (
	"github.com/derektamsen/cdk8ssharedapp"
)

func main() {
	clusters := &cdk8ssharedapp.K8sClusters{
		Clusters: &[]cdk8ssharedapp.ClusterProps{
			{ClusterName: "stage", Image: "nginx:latest"},
			{ClusterName: "prod", Image: "nginx:1.19.10"},
		},
	}

	cdk8ssharedapp.NewApp(clusters)
}
