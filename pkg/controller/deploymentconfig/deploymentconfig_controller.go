package deploymentconfig

import (
	"context"
	"fmt"
	"time"

	appsv1 "underThedome-operator/pkg/apis/apps/v1"

	"underThedome-operator/pkg/controller/underthedome"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	//"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
	//"github.com/davecgh/go-spew/spew"
        "net/url"
)

var log = logf.Log.WithName("controller_deploymentconfig")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new DeploymentConfig Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileDeploymentConfig{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("deploymentconfig-controller", mgr, controller.Options{Reconciler: r,MaxConcurrentReconciles: 20})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource DeploymentConfig
	err = c.Watch(&source.Kind{Type: &appsv1.DeploymentConfig{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner DeploymentConfig
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appsv1.DeploymentConfig{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileDeploymentConfig implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileDeploymentConfig{}

// ReconcileDeploymentConfig reconciles a DeploymentConfig object
type ReconcileDeploymentConfig struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a DeploymentConfig object and makes changes based on the state read
// and what is in the DeploymentConfig.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileDeploymentConfig) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling DeploymentConfig")

	wait_inited()
	
	
	if checkNamespace(request.Namespace) == true {
	// Fetch the DeploymentConfig instance
	instance := &appsv1.DeploymentConfig{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	reqLogger.Info("DeploymentConfig event on ", "DeploymentConfig.Namespace", instance.Namespace, "DeploymentConfig.Name", instance.Name)
	for _, cont := range instance.Spec.Template.Spec.Containers {
		
		fmt.Printf("Container Name %s Container\n",cont.Name)
		// Imagestramse check 
		if checkImage(cont.Image) != true {
			fmt.Printf("Invalid Image Jeiling Container %s\n",cont.Name)
			instance.Spec.Replicas=0
                        annotations:=instance.ObjectMeta.GetAnnotations()
                        if( annotations == nil ) {
                                an:=make(map[string]string)
                                an["under.the.dome/jailed"]="true"
                                instance.ObjectMeta.SetAnnotations(an)
                        } else {
                                instance.ObjectMeta.Annotations["under.the.dome/jailed"]="true"
                        }
                	_ = r.client.Update(context.TODO(),instance)

		} else {
			fmt.Printf("Image is Valid\n")
		}
		
	}
	} else {
		fmt.Printf("Namespace %s nothing to do\n",request.Namespace)
	}
	return reconcile.Result{}, nil
}

func checkImage(image string) bool {
		ur:="https://"+image
		i, err:=  url.Parse(ur)	
		if err != nil {
			fmt.Printf("Error is %s\n",err)
			return false
		}
                host:=""
                if ( i.Port() != "") {
                        host=i.Hostname()+":"+i.Port()
                } else {
                        host=i.Hostname()
                }
		fmt.Printf("Registry is on registry %s checking validity\n",host)
		return checkRepository(host)
}

func checkNamespace(nameSpace string) bool {
	for _ , v := range underthedome.UnderTheDome_instance.Spec.Namespaces {
		if ( v == nameSpace ) {
			return true
		}
	} 	
	return false
}

func checkRepository(repo string) bool {
	for _ , v := range underthedome.UnderTheDome_instance.Spec.Repositories {
		if ( v == repo ) {
			return true
		}
	} 	
	return false
}

func wait_inited() {
	for underthedome.Inited == false {
		time.Sleep(1 * time.Second)
	}
}
