package v1

import (
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        opeanshift "github.com/openshift/api/apps/v1"
	openshiftv1image "github.com/openshift/api/image/v1"
)
type DeploymentConfig = opeanshift.DeploymentConfig
type DeploymentConfigList = opeanshift.DeploymentConfigList
type ImageStream = openshiftv1image.ImageStream
type ImageStreamList = openshiftv1image.ImageStreamList

func init() {
	SchemeBuilder.Register(&opeanshift.DeploymentConfig{}, &opeanshift.DeploymentConfigList{})
}
