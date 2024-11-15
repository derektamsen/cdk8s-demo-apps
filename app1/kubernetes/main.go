package main

import (
	"github.com/derektamsen/cdk8ssharedapp"
)

func main() {
	appConfig := &cdk8ssharedapp.AppConfig{
		Name:      "app1",
		Namespace: "app1",
	}

	clusters := &cdk8ssharedapp.K8sClusters{
		Clusters: &[]cdk8ssharedapp.ClusterProps{
			{ClusterName: "stage-cluster-01", Image: "ruby:latest"},
			{ClusterName: "prod-cluster-01", Image: "ruby:3.1-alpine"},
		},
	}

	cdk8ssharedapp.NewApp(appConfig, clusters)
}
