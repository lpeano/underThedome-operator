package underthedome

import (
	"context"
	"fmt"
	"os"

	underthedomev1 "underThedome-operator/pkg/apis/underthedome/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
        //"k8s.io/client-go/tools/clientcmd"
	//openshiftv1 "github.com/openshift/api/apps/v1"
	//openshiftv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
        "k8s.io/client-go/util/workqueue"
	"github.com/davecgh/go-spew/spew"
)

/*
type GenericConfig struct {
       openshiftv1.DeploymentConfig{}
       DeploymentConfigInterface{}
}*/

var UnderTheDome_instance =  &underthedomev1.Underthedome{}
var Inited bool = false


var log = logf.Log.WithName("controller_underthedome")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Underthedome Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileUnderthedome{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("underthedome-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Underthedome
	//err = c.Watch(&source.Kind{Type: &underthedomev1.Underthedome{}}, &handler.EnqueueRequestForObject{})
	err = c.Watch(&source.Kind{Type: &underthedomev1.Underthedome{}}, handler.Funcs{
			CreateFunc: func(e event.CreateEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.Meta.GetName(),
					Namespace: e.Meta.GetNamespace(),
				}})
			},	
			UpdateFunc: func(e event.UpdateEvent, q workqueue.RateLimitingInterface) {
				os.Exit(143)
			},
			DeleteFunc: func(e event.DeleteEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.Meta.GetName(),
					Namespace: e.Meta.GetNamespace(),
				}})
			},
			GenericFunc: func(e event.GenericEvent, q workqueue.RateLimitingInterface) {
				q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
					Name:      e.Meta.GetName(),
					Namespace: e.Meta.GetNamespace(),
				}})
			},	

		})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileUnderthedome implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileUnderthedome{}

// ReconcileUnderthedome reconciles a Underthedome object
type ReconcileUnderthedome struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Underthedome object and makes changes based on the state read
// and what is in the Underthedome.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileUnderthedome) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Underthedome")
        spew.Dump(request)

	// Fetch the Underthedome UnderTheDome_instance
	err_ud := r.client.Get(context.TODO(),  types.NamespacedName{Namespace:request.Namespace,Name:request.Name}, UnderTheDome_instance)
	if err_ud != nil {
		if errors.IsNotFound(err_ud) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			reqLogger.Info("error NotFound - Requeueing request ", request.NamespacedName, "erroe",err_ud)
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		reqLogger.Info("error - Requeueing request ", request.NamespacedName, "erroe",err_ud)
		return reconcile.Result{}, err_ud
	}
	fmt.Printf("UnderTheDome_instance.Spec.Namespaces:%v",UnderTheDome_instance.Spec.Namespaces)
	
	Inited = true
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *underthedomev1.Underthedome) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}


func MatchNamespaces(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
