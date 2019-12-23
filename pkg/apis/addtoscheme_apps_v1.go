package apis

import v1 "underThedome-operator/pkg/apis/underthedome/v1"
import openshiftv1image "github.com/openshift/api/image/v1"

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1.SchemeBuilder.AddToScheme)
	AddToSchemes = append(AddToSchemes, openshiftv1image.SchemeBuilder.AddToScheme)
}
