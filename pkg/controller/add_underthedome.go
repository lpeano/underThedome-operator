package controller

import (
	"underThedome-operator/pkg/controller/underthedome"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, underthedome.Add)
}
