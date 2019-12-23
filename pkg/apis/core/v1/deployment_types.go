package v1

import (
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cv1 "k8s.io/api/apps/v1"
)

type Deployment = cv1.Deployment
type DeploymentList = cv1.DeploymentList

func init() {
	SchemeBuilder.Register(&cv1.Deployment{}, &cv1.DeploymentList{})
}
