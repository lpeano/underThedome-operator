# underThedome-operator
Checke Image Registry validiti
The "underthedome" operator implements the control of the goodness of the registry of origin of the deployment, deploymentconfig and statefulset images.

If the control is invalid (based on the operator configuration), scale the resources involved to zero and note the resource as "jailed".

The example configuration is shown below.


```
NAME                   AGE
example-underthedome   26d
[lpeano@oclab ~]$ oc get underthedomes -o yaml --export
apiVersion: v1
items:
- apiVersion: underthedome.extentsions.io/v1
  kind: Underthedome
  metadata:
    creationTimestamp: 2019-12-18T01:56:29Z
    generation: 80
    name: example-underthedome
    namespace: underthedome
    resourceVersion: "1389672"
    selfLink: /apis/underthedome.extentsions.io/v1/namespaces/underthedome/underthedomes/example-underthedome
    uid: 9b9391fa-2139-11ea-b0de-08002730905e
  spec:
    namespaces:
    - prova2
    - peano
    - thresholds
    - operator
    repositories:
    - 172.30.1.1:5000
    - k8s.gcr.io
    watchnamespace: underthedome
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""
```

Where 'namespaces' is a list of namespaces to be monitored and repositories is a list oc valid refistries.
