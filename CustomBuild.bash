export GO111MODULE=on
export GOFLAGE="-v 5"
export GOPATH=$(go env GOPATH)
export GOROOT=/usr/lib/golang
export WATCH_NAMESPACE=""

export GOROOT=/usr/lib/golang
export GO111MODULE=on
operator-sdk generate k8s && operator-sdk build underthedome-operator:v0.0.1 --verbose && (
oc login -u system -p admin
oc project underthedome
docker save underthedome-operator:v0.0.1  -o underthedome-operator_v0.0.1
skopeo copy --dest-tls-verify=false --dest-creds  system:$(oc whoami -t)  docker-archive:./underthedome-operator_v0.0.1:underthedome-operator:v0.0.1  docker://172.30.1.1:5000/underthedome/underthedome-operator:v0.0.1
)
rm -f underthedome-operator_v0.0.1
