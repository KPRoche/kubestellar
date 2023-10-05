# Automate the kubestellar core side of the extended example
# Make sure "echo $HOSTNAME" returns the FQDN of your host machine
# execute as ". example_1.sh $HOSTNAME:[port]"
# make sure the yaml files common, special, edgeplacement and edgeplacement-s are in the folder from which you call this script

kubectl kubestellar deploy --external-endpoint $1
# this takes a while to fully deploy and start running
echo "kubestellar started with endpoint " $1
echo "Wait a few minutes for all the containers to come up to speed before continuing"
read -p "Press enter key to resume ..."

echo "requesting kubestellar kubeconfig as ks.config"
kubectl kubestellar get-external-kubeconfig -o ks.config
export KUBECONFIG=$(pwd)/ks.config
ls -l $KUBECONFIG
kubectl ws tree
kubectl ws root
kubectl ws create imw-1 
kubectl ws root:imw-1
kubectl kubestellar ensure location florin  loc-name=florin  env=prod
kubectl kubestellar ensure location guilder loc-name=guilder env=prod extended=si
kubectl get Locations
kubectl get Workspaces
kubectl ws root:espw
kubectl get Workspaces
kubectl get Workspace -o "custom-columns=NAME:.metadata.name,SYNCTARGET:.metadata.annotations['edge\.kubestellar\.io/sync-target-name'],CLUSTER:.spec.cluster"
kubectl get Workspace -o "custom-columns=NAME:.metadata.name,SYNCTARGET:.metadata.annotations['edge\.kubestellar\.io/sync-target-name'],CLUSTER:.spec.cluster" > workspaces.txt
GUILDER_WS=$(kubectl get Workspace -o json | jq -r '.items | .[] | .metadata | select(.annotations ["edge.kubestellar.io/sync-target-name"] == "guilder") | .name')
echo The guilder mailbox workspace name is $GUILDER_WS
FLORIN_WS=$(kubectl get Workspace -o json | jq -r '.items | .[] | .metadata | select(.annotations ["edge.kubestellar.io/sync-target-name"] == "florin") | .name')
echo The florin mailbox workspace name is $FLORIN_WS
kubectl kubestellar prep-for-syncer --imw root:imw-1 florin
kubectl kubestellar prep-for-syncer --imw root:imw-1 guilder
kubectl ws root
kubectl ws create my-org --enter
kubectl kubestellar ensure wmw wmw-c
kubectl ws tree
kubectl apply -f common.yaml 
kubectl apply -f edgeplacement.yaml 
kubectl get ns
kubectl get ReplicaSets -A
kubectl get deployments,ReplicaSets -A
# now do the special case
kubectl ws root:my-org
kubectl kubestellar ensure wmw wmw-s
kubectl ws tree
kubectl apply -f special.yaml 
kubectl apply -f edgeplacement-s.yaml 
kubectl get ns
kubectl get ReplicaSets -A
kubectl get deployments,ReplicaSets -A