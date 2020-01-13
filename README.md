# underThedome-operator
Checke Image Registry validiti
The "underthedome" operator implements the control of the goodness of the registry of origin of the deployment, deploymentconfig and statefulset images.

If the control is invalid (based on the operator configuration), scale the resources involved to zero and note the resource as "jailed".
