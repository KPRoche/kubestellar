# KubeStellar Hub chart usage

This documents explains how to use KubeStellar Hub chart to deploy a new instance of KubeStellar
with a choice of user-defined KubeFlex Control Planes (CPs).

The information provided is specific for the following release:

```shell
export KUBESTELLAR_VERSION=0.23.0-alpha.2
```

## Pre-requisites

To install the Helm chart the only requirement is [Helm](https://helm.sh/).
However, additional executables may be required to create/manage the cluster(s) (_e.g._, Kind and kubectl),
to join Workload Execution Clusters (WEKs) (_e.g._, clusteradm),
and to interact with Control Planes (_e.g._, kubectl), _etc_.
For such purpose, a full list of executable that may be required can be fund [here](./pre-reqs.md).

The setup of KubStellar via the Hub chart requires the existance of at least one cluster
to be used for the deployment of the chart.

This can be:

1. A local **Kind** cluster with an ingress with SSL passthrough and a mapping to host port 9443

This option is particularly useful for first time users or users that would like to have a local deployment.

It is important to note that, if a Kind cluster name different from `kubeflex` is used, then its Control Plane name must be also be referenced during the Helm chart installation by setting `--set "kubeflex-controller.hostContainer=<control-plane-name>"`. The Control Plane name can be obtained by using `docker ps` to find the name of the container running the Kind cluster used for deploying KubeStellar Hub chart, _e.g._ `kubeflex-control-plane`.

If a host port number different from the expected 9443 is used for the Kind cluster, then the same port number must be specified during the chart installation by adding the following argument `--set "kubeflex-controller.externalPort=<port>"`.

By default the KubeStelalr Hub chart uses a test domanin `localtest.me`, which is ok for testing on a single host machine. However, scenarios that span more than one machine, it may be useful to set `--set "kubeflex-controller.domanin=<domain>"` to a more appropriate `<domain>` that can be reached from Workload Execution CLusters (WECs).

For convenience, a new local Kind cluster that satisfies the requirements for KubeStellar setup
and that can be used to exercises the [examples](./examples.md) can be created with the following command:

```shell
sudo su <<EOF
bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/v${KUBESTELLAR_VERSION}/scripts/create-kind-cluster-with-SSL-passthrough.sh) --name kubeflex --port 9443
EOF
```

2. A local **k3s** cluster with an ingress with SSL passthrough and a mapping to host port 9443

... Under onstruction ...

3. An **OpenShift** cluster

When using this option, one is required to explicitely set the `isOpenShift` variable to `true` by including `--set "kubeflex-operator.isOpenShift=true"` in the Helm chart installation command.


## KubeStellar Hub Chart values

The KubeStellar chart makes available to the user several values that may be used to customize its installation into an existing cluster:

```yaml
# KubeFlex override values
kubeflex-operator:
  install: true # enable/disable the installation of KubeFlex by the chart (default: true)
  installPostgreSQL: true # enable/disable the installation of the appropriate version of PostgreSQL required by KubeFlex (default: true)
  isOpenShift: false # set this variable to true when installing the chart in an OpenShift cluster (default: false)
  # Kind cluster specific settings:
  domain: localtest.me # used to set the Control Planes DNS in a Kind cluster installation (default: localtest.me)
  externalPort: 9443 # used to set the port to access the Control Planes API (default: 9443)
  hostContainer: kubeflex-control-plane # used to set the name of cluster control plane (default: kubeflex-control-plane, which corresponds to a Kind cluster with name kubeflex)

# Determine if the Post Create Hooks should be installed by the chart
InstallPCHs: true

# List the Inventory and Transport Spaces (ITSes) to be created by the chart
# Each ITS consists of a mandatory unique name and an optional type, which could be either host|vcluster (default to vcluster, if not specified)
ITSes: # ==> installs ocm + ocm-status-addon

# List the Workload Description Spaces (WDSes) to be created by the chart
# Each WDS consists of a mandatory unique name and several optional parameters:
# - type: host or k8s (default to k8s, if not specified)
# - APIGroups: a comma separated list of APIGroups
# - ITSName: the name of the ITS control plane to be used by the WDS. Note that the ITSName MUST be specified if more than one ITS exists.
WDSes: # ==> installs kubestellar + ocm-transport-plugin
```

The first section of the `values.yaml` file refers to parameters that are specific to the KubeFlex instllation, see [here](https://github.com/kubestellar/kubeflex/blob/main/docs/users.md) for more information.

In particular:
- `kubeflex-operator.install` accepts a boolean value to enable/disable the installation of KubeFlex into the cluster by the chart
- `kubeflex-operator.isOpenShift` must be set to true by the user when installing the chart into a OpenShift cluster

By default, the chart will install the KubeFlex and its PostgreSQL dependency.

The second section allows a user of the chart to determine if Post Create Hooks (PCHs) needed for creating ITSes and WDSes control planes should be instlalled by the chart. By default `InstallPCHs` is set to `true` to enable the instllation of the PCHs, however one may want to set this value to `false` when installing multiple copies of the chart to avoid conflicts. A single copy of the PCHs is required and allowed per cluster.

The third section of the `values.yaml` file allows one to create a list of Inventory and Transport Spaces (ITSes). By default, this list is empty and no ITS will be created by the chart. A list of ITSes can be specified using the following format:

```yaml
ITSes: # all the CPs in this list will execute the its.yaml PCH
  - name: <its1>          # mandatory name of the control plane
    type: <vlcuster|host> # optional type of control plane host or vcluster (default to vcluster, if not specified)
  - name: <its2>          # mandatory name of the control plane
    type: <vlcuster|host> # optional type of control plane host or vcluster (default to vcluster, if not specified)
  ...
```

where `name` must specify a unique name of the control plane and the optional `type` can be either vlcuster (default) or host, see [here](https://github.com/kubestellar/kubeflex/blob/main/docs/users.md) for more information.

The fourth section of the `values.yaml` file allows one to create a list of Workload Description Spaces (WDSes). By default, this list is empty and no WDS will be created by the chart. A list of WDSes can be specified using the following format:

```yaml
WDSes: # all the CPs in this list will execute the wds.yaml PCH
  - name: <wds1>     # mandatory name of the control plane
    type: <host|k8s> # optional type of control plane host or k8s (default to k8s, if not specified)
    APIGroups: ""    # optional string holding a comma-separated list of APIGroups
    ITSName: <its1>  # optional name of the ITS control plane, this MUST be specified if more than one ITS exists at the moment the WDS PCH starts
  - name: <wds2>     # mandatory name of the control plane
    type: <host|k8s> # optional type of control plane host or k8s (default to k8s, if not specified)
    APIGroups: ""    # optional string holding a comma-separated list of APIGroups
    ITSName: <its2>  # optional name of the ITS control plane, this MUST be specified if more than one ITS exists at the moment the WDS PCH starts
  ...
```

where `name` must specify a unique name of the control plane (not that this must be unique among both ITSes and WDSes), the optional `type` can be either k8s (default) or host, see [here](https://github.com/kubestellar/kubeflex/blob/main/docs/users.md) for more information, the optional `APIGroups` provides a list of APIGroups, see [here](https://docs.kubestellar.io/release-0.22.0/direct/examples/#scenario-2-using-the-hosting-cluster-as-wds-to-deploy-a-custom-resource) for more information, and `ITSName` specify the ITS connected to the new WDS being created (this parameter MUST specified only if more that one ITS exists in the cluster, if no value is specified and only one ITS exists in the cluster, then it will automatically selected).

## KubeStellar Hub Chart usage

A specific version of the KubeStellar hub chart can be simply installed in an existing cluster using the following command:

```shell
helm upgrade --install ks-hub oci://ghcr.io/kubestellar/kubestellar/hub-chart --version $KUBESTELLAR_VERSION
```

The above command will install KubeFlex and the Post Create Hooks, but no Control Planes.
Please remeber to add `--set "kubeflex-operator.isOpenShift=true"`, when installing into an OpenShift cluster.

User defined control planes can be added using additional value files of `--set` arguments, _e.g._:

- add a single ITS named its1 of default vcluster type: `--set-json='ITSes=[{"name":"its1"}]'`
- add two ITSes named its1 and its2 of of type vlcuster and host, respectevely: `--set-json='ITSes=[{"name":"its1"},{"name":"its2","type":"host"}]'`
- add a single WDS named wds1 of default k8s type: `--set-json='WDSes=[{"name":"wds1"}]'`

A KubeStellar Hub installation compatible with the common setup suitable for Common Setup described in the [examples](./examples.md) could be achieved with the following command:

```shell
helm upgrade --install ks-hub oci://ghcr.io/kubestellar/kubestellar/hub-chart --version $KUBESTELLAR_VERSION \
  --set-json='ITSes=[{"name":"its1"}]' \
  --set-json='WDSes=[{"name":"wds1"}]'
```

After the initial instllation is completed, there are two main ways to install additional control planes (_e.g._, create a second `wds2` WDS):

1. Upgrade the initial chart. This choice requires to relist the existing control planes, which would otherwise be deleted:

```shell
helm upgrade --install ks-hub oci://ghcr.io/kubestellar/kubestellar/hub-chart --version $KUBESTELLAR_VERSION \
  --set-json='ITSes=[{"name":"its1"}]' \
  --set-json='WDSes=[{"name":"wds1"},{"name":"wds2"}]'
```

2. Install a new chart with a different name. This choice does not requires to relist the existing control planes, but requies to disable the reinstallation of KubeFlex and PCHs:

```shell
helm upgrade --install add-wds2 oci://ghcr.io/kubestellar/kubestellar/hub-chart --version $KUBESTELLAR_VERSION \
  --set='kubeflex-operator.install=false,InstallPCHs=false' \
  --set-json='WDSes=[{name":"wds2"}]'
```

## Obtaining the Control Planes kubeconfigs

The following code can be used to wait for the Ready state of all the KubeFlex Control Planes in the hub cluster:

```shell
echo "Waiting for all Control Planes to be ready..."
for cpname in `kubectl get controlplane -o name`; do
  cpname=${cpname##*/}
  while [[ $(kubectl get cp $cpname -o 'jsonpath={..status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do
    echo "Waiting for Control Plane $cpname..."
    sleep 5
  done
  echo $cpname
done
```

The following code can be used to obtain the kubeconfig of all the KubeFlex Control Planes in the hub cluster:

```shell
echo "Getting the kubeconfig of all Control Planes..."
for cpname in `kubectl get controlplane -o name`; do
  cpname=${cpname##*/}
  echo "Getting the kubeconfig of Control Planes $cpname ==> /tmp/kubeconfig-$cpname..."
  if [[ "$(kubectl get controlplane $cpname -o=jsonpath='{.spec.type}')" == "host" ]] ; then
    kubectl config view --minify --flatten > "kubeconfig-$cpname"
  else
    kubectl get secret $(kubectl get controlplane $cpname -o=jsonpath='{.status.secretRef.name}') \
      -n $(kubectl get controlplane $cpname -o=jsonpath='{.status.secretRef.namespace}') \
      -o=jsonpath="{.data.$(kubectl get controlplane $cpname -o=jsonpath='{.status.secretRef.key}')}" \
      | base64 -d > "kubeconfig-$cpname"
  fi
  kubectl --kubeconfig "kubeconfig-$cpname" config rename-context $(kubectl --kubeconfig "kubeconfig-$cpname" config current-context) $cpname 2> /dev/null
done
```

The code above will saves the kubeconfig of a control plane `$cpname` to a corresponding file `kubeconfig-$cpname` in the local folder.
The content of a Control Plane `$cpname` can be accessed by specifing its kubeconfig:

```shell
kubectl --kubeconfig "kubeconfig-$cpname" ...
```

The individual kubeconfigs can also be merged as contexts of the current `~/.kube/config` with the following command:

```shell
KUBECONFIG=~/.kube/config:$(find . -maxdepth 1 -type f -name 'kubeconfig-*' | tr '\n' ':') kubectl config view --flatten > /temp/kubeconfig-merged
cp /temp/kubeconfig-merged ~/.kube/config
```

Afterwards the content of a Control Plane `$cpname` can be accessed by specifing its context:

```shell
kubectl --context "$cpname" ...
```

## Uninstalling the KubeStellar Hub chart

The chart can be uninsitalled using the command:

```shell
helm uninstall ks-hub
```

This will remove KubeFlex, Post Create Hooks, and all KubeFlex Control Planes that were created by the chart.