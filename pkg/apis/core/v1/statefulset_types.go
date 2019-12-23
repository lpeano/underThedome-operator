package v1

import (
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cv1 "k8s.io/api/apps/v1"
)

type StatefulSet = cv1.StatefulSet 
type StatefulSetList = cv1.StatefulSetList

func init() {
        SchemeBuilder.Register(&cv1.StatefulSet{}, &cv1.StatefulSetList{})
}

