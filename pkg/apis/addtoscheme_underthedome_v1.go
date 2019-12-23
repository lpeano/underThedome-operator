package apis

import (
	v1 "underThedome-operator/pkg/apis/underthedome/v1"
        openshiftv1 "github.com/openshift/api/apps/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
	AddToSchemes = append(AddToSchemes, openshiftv1.SchemeBuilder.AddToScheme)
}
