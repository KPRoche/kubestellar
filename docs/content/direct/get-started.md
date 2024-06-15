# KubeStellar Quickstart Setup

This Quick Start is based on Scenario 1 of our [examples](examples.md).
In a nutshell, you will:

  1. Before you begin, prepare your system (get the prerequisites)
  2. Create the Kubestellar core components on a cluster
  3. Commission a workload to a WEC

---
## Before You Begin


{%
    include-markdown "pre-reqs.md"
    end="For Running"
    heading-offset=2
%}

---

## Create the KubeStellar Core components

Use our helm chart to set up the main core and establish its initial state using our helm chart:

### Set the Version appropriately as an environment variable

```shell
export KUBESTELLAR_VERSION=0.23.0-alpha.4
```
### Use the Helm chart  to deploy the KubeStellar Core to a Kind, K3s, or OpenShift cluster:

#### Using Kind

For convenience, a new local **Kind** cluster that satisfies the requirements for KubeStellar setup
and that can be used to commission the quickstart workload can be created with the following command:

```shell
bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/v0.23.0-alpha.4/scripts/create-kind-cluster-with-SSL-passthrough.sh) --name kubeflex --port 9443
```
After the cluster is created, deploy the Kubestellar Core installation on it with the helm chart command

```shell
helm upgrade --install ks-core oci://ghcr.io/kubestellar/kubestellar/core-chart --version $KUBESTELLAR_VERSION \
  --set-json='ITSes=[{"name":"its1"}]' \
  --set-json='WDSes=[{"name":"wds1"}]'
```

#### using K3S

A new local **k3s** cluster that satisfies the requirements for KubeStellar setup
and that can be used to commission the quickstart workload can be created with the following command:

```shell
bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/v0.23.0-alpha.4/scripts/create-k3s-cluster-with-SSL-passthrough.sh) --port 9443
```
After the cluster is created, deploy the Kubestellar Core installation on it with the helm chart command

```shell
helm upgrade --install ks-core oci://ghcr.io/kubestellar/kubestellar/core-chart --version $KUBESTELLAR_VERSION \
  --set-json='ITSes=[{"name":"its1"}]' \
  --set-json='WDSes=[{"name":"wds1"}]'
```

#### Using OpenShift

When using this option, one is required to explicitely set the `isOpenShift` variable to `true` by including `--set "kubeflex-operator.isOpenShift=true"` in the Helm chart installation command.

After the cluster is created, deploy the Kubestellar Core installation on it with the helm chart command

```shell
helm upgrade --install ks-core oci://ghcr.io/kubestellar/kubestellar/core-chart --version $KUBESTELLAR_VERSION \
  --set "kubeflex-operator.isOpenShift=true" \ 
  --set-json='ITSes=[{"name":"its1"}]' \
  --set-json='WDSes=[{"name":"wds1"}]'
```

Once you have done this, you should have the KubeStellar core components plus the required workload definition space and inventory and transport space control planes running on your cluster.

---

## Define, bind and commission a workload on a WEC

  {%
    include-markdown "example-wecs.md"
    heading-offset=2
  %}

