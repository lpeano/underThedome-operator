package underthedome

import (
	"context"
	"fmt"
	//"os"
	"net/url"
	"strconv"

        corev1api "underThedome-operator/pkg/apis/core/v1"
        appsv1 "underThedome-operator/pkg/apis/apps/v1"
	underthedomev1 "underThedome-operator/pkg/apis/underthedome/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	//"sigs.k8s.io/controller-runtime/pkg/event"
        //"k8s.io/client-go/util/workqueue"
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
	err = c.Watch(&source.Kind{Type: &underthedomev1.Underthedome{}}, &handler.EnqueueRequestForObject{})
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
	
	CheckAllNS(request,r)
	Inited = true
	return reconcile.Result{}, nil
}

func CheckAllNS(request reconcile.Request,r *ReconcileUnderthedome){
	for _, v := range  UnderTheDome_instance.Spec.Namespaces {
		fmt.Printf("Checking DeploymentConfig for: %s\n", v)	
		GetDeploymentconfig(v,r)
		fmt.Printf("Checking Deployment for: %s\n", v)	
		GetDeployment(v,r)
		fmt.Printf("Checking Statefulset for: %s\n", v)	
		GetStatefulSet(v,r)
	}

}

func GetDeployment(namespace string,r *ReconcileUnderthedome){
	instList:=&corev1api.DeploymentList{}
	err := r.client.List(context.TODO(),&client.ListOptions{Namespace:namespace},instList)
	if err != nil {
		fmt.Printf("No deployment found\n")
	} else {
        	//err := r.client.Get(context.TODO(),  types.NamespacedName{Namespace:namespace,Name:name}, instance)
		for k , d := range instList.Items {
			fmt.Printf("Item N° %d\n",  k)
			if CheckContainers(d.Spec.Template.Spec.Containers) != true {
				fmt.Printf("StatefulSet - Namespace: %s Name: %s Invalid Container found\n", namespace, d.ObjectMeta.Name)
                        	annotations:=d.ObjectMeta.GetAnnotations()
                        	if( annotations == nil ) {
                                	an:=make(map[string]string)
                                	an["under.the.dome/jailed"]="true"
					an["under.the.dome/replicas"]=strconv.FormatInt(int64(*d.Spec.Replicas), 10)               
                                	d.ObjectMeta.SetAnnotations(an)
                        	} else {
                                	d.ObjectMeta.Annotations["under.the.dome/jailed"]="true"
					d.ObjectMeta.Annotations["under.the.dome/replicas"]=strconv.FormatInt(int64(*d.Spec.Replicas), 10)
                        	}
	                        x:=int32(0)
       	                        d.Spec.Replicas=&x
                        	_ = r.client.Update(context.TODO(),&d)

				
			} else {

				fmt.Printf("StatefulSet - Namespace: %s Name: %s Invalid Container found\n", namespace, d.ObjectMeta.Name)
                        	annotations:=d.ObjectMeta.GetAnnotations()
                        	if( annotations == nil ) {
                                	an:=make(map[string]string)
                                	an["under.the.dome/jailed"]="false"
                                	d.ObjectMeta.SetAnnotations(an)
                        	} else {
                                	d.ObjectMeta.Annotations["under.the.dome/jailed"]="false"
                        	}
				i32,_:=strconv.ParseInt(d.ObjectMeta.Annotations["under.the.dome/replicas"],10,32)
				x:=int32(i32)
				d.Spec.Replicas=&x
                        	_ = r.client.Update(context.TODO(),&d)
			} 
		}
		fmt.Printf("\nError is %v\n",err)
		//os.Exit(0)
	}
}

func GetDeploymentconfig(namespace string,r *ReconcileUnderthedome){
	instList:=&appsv1.DeploymentConfigList{}
	err := r.client.List(context.TODO(),&client.ListOptions{Namespace:namespace},instList)
	if err != nil {
		fmt.Printf("No deployment found\n")
	} else {
		for k , d := range instList.Items {
			fmt.Printf("Item N° %d\n",  k)
			if CheckContainers(d.Spec.Template.Spec.Containers) != true {
				fmt.Printf("StatefulSet - Namespace: %s Name: %s Invalid Container found\n", namespace, d.ObjectMeta.Name)
                        	annotations:=d.ObjectMeta.GetAnnotations()
                        	if( annotations == nil ) {
                                	an:=make(map[string]string)
                                	an["under.the.dome/jailed"]="true"
					an["under.the.dome/replicas"]=strconv.FormatInt(int64(d.Spec.Replicas), 10)       
                                	d.ObjectMeta.SetAnnotations(an)
                        	} else {
                                	d.ObjectMeta.Annotations["under.the.dome/jailed"]="true"
                                	d.ObjectMeta.Annotations["under.the.dome/replicas"]=strconv.FormatInt(int64(d.Spec.Replicas), 10)
                        	}
       	                        d.Spec.Replicas=0
                        	_ = r.client.Update(context.TODO(),&d)

				
			} else {

				fmt.Printf("StatefulSet - Namespace: %s Name: %s Invalid Container found\n", namespace, d.ObjectMeta.Name)
                        	annotations:=d.ObjectMeta.GetAnnotations()
                        	if( annotations == nil ) {
                                	an:=make(map[string]string)
                                	an["under.the.dome/jailed"]="false"
                                	d.ObjectMeta.SetAnnotations(an)
                        	} else {
                                	d.ObjectMeta.Annotations["under.the.dome/jailed"]="false"
					i32,_:=strconv.ParseInt(d.ObjectMeta.Annotations["under.the.dome/replicas"],10,32)
					x:=int32(i32)
					d.Spec.Replicas=x
                        	}
                        	_ = r.client.Update(context.TODO(),&d)
			} 
		}
		fmt.Printf("\nError is %v\n",err)
	}
}

func GetStatefulSet(namespace string,r *ReconcileUnderthedome){
	instList:=&corev1api.StatefulSetList{}
	err := r.client.List(context.TODO(),&client.ListOptions{Namespace:namespace},instList)
	if err != nil {
		fmt.Printf("No deployment found\n")
	} else {
        	//err := r.client.Get(context.TODO(),  types.NamespacedName{Namespace:namespace,Name:name}, instance)
		for k , d := range instList.Items {
			fmt.Printf("Item N° %d\n",  k)
			if CheckContainers(d.Spec.Template.Spec.Containers) != true {
				fmt.Printf("StatefulSet - Namespace: %s Name: %s Invalid Container found\n", namespace, d.ObjectMeta.Name)
                        	annotations:=d.ObjectMeta.GetAnnotations()
                        	if( annotations == nil ) {
                                	an:=make(map[string]string)
                                	an["under.the.dome/jailed"]="true"
					an["under.the.dome/replicas"]=strconv.FormatInt(int64(*d.Spec.Replicas), 10)               
                                	d.ObjectMeta.SetAnnotations(an)
                        	} else {
                                	d.ObjectMeta.Annotations["under.the.dome/jailed"]="true"
					d.ObjectMeta.Annotations["under.the.dome/replicas"]=strconv.FormatInt(int64(*d.Spec.Replicas), 10)
                        	}
	                        x:=int32(0)
       	                        d.Spec.Replicas=&x
                        	_ = r.client.Update(context.TODO(),&d)

				
			} else {

				fmt.Printf("StatefulSet - Namespace: %s Name: %s Invalid Container found\n", namespace, d.ObjectMeta.Name)
                        	annotations:=d.ObjectMeta.GetAnnotations()
                        	if( annotations == nil ) {
                                	an:=make(map[string]string)
                                	an["under.the.dome/jailed"]="false"
                                	d.ObjectMeta.SetAnnotations(an)
                        	} else {
                                	d.ObjectMeta.Annotations["under.the.dome/jailed"]="false"
                        	}
				i32,_:=strconv.ParseInt(d.ObjectMeta.Annotations["under.the.dome/replicas"],10,32)
				x:=int32(i32)
				d.Spec.Replicas=&x
                        	_ = r.client.Update(context.TODO(),&d)
			} 
		}
		fmt.Printf("\nError is %v\n",err)
	}
}

func CheckContainers(Containers []corev1.Container) bool {
	for _, c := range Containers{
		fmt.Printf("Container %s checking\n", c.Name)
		if checkImage(c.Image) != true {
			return false
		}		
	}
	return true
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

func checkRepository(repo string) bool {
        for _ , v := range UnderTheDome_instance.Spec.Repositories {
                if ( v == repo ) {
                        return true
                }
        }
        return false
}

