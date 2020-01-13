# underThedome-operator
Checke Image Registry validiti
The "underthedome" operator implements the control of the goodness of the registry of origin of the deployment, deploymentconfig and statefulset images.

If the control is invalid (based on the operator configuration), scale the resources involved to zero and note the resource as "jailed".

The example configuration is shown below.


>apiVersion: v1
>items:
>- apiVersion: underthedome.extentsions.io/v1
>  kind: Underthedome
>  metadata:
>    selfLink: /apis/underthedome.extentsions.io/v1/namespaces/underthedome/underthedomes/example-underthedome
>    uid: 9b9391fa-2139-11ea-b0de-08002730905e
>  spec:
>    namespaces:
>    - prova2
>    - peano
>    - thresholds
>    - operator
>    repositories:
>    - 172.30.1.1:5000
>    - k8s.gcr.io
>    watchnamespace: underthedome
>kind: List
>metadata:
>  resourceVersion: ""
>  selfLink: ""
