# General overview of KubeStellar implemenation

KubeStellar operates with a multiple-cluster [architecture](readme.md):

* The KubeStellar core components (referred to as a "hub" in some of the example documents)
* One or more Workload Execution Clusters (WECs) to which the hub deploys workloads
_(there is support for the special case in which the workload deploys to the same cluster as the hub)_


The general flow for setting up and using KubeStellar is in 3 steps:

1. Set Up the KubeStellar Hub
2. Set Up and Register the Workload Execution Clusters
3. Define and Deploy workloads to the WECs

Given the multitude of infrastructure configurations possible, the details of a particular installation will vary.

We have developed a [common setup](direct/common-setup-intro.md) for our examples as a useful starting point. You can use a [helm chart](direct/common-setup-hub-chart.md) to automate the process of setting up the core components (hub), or you may [work through the steps individually](direct/common-setup-step-by-step.md) to more completely customize the installation.

<!--
* Set up infrastructure to host the hub and workload clusters
* Install prerequisite software to do the setup
* Set up the KubeStellar core components (hub) cluster(s)
* Set up Workload Execution Cluster(s)
* Register WECs with the hub
* Define workloads for deployment
* Deploy workloads
* Confirm/monitor status of workload(s)
* Redefine workloads as necessary (Updates/Undeploys/Redeploys workload on WECs)
-->
<!-- ## Prereqs

### Set up your core and workload cluster infrastructure

### install appropriate software there

## Set up the KubeStellar Core components

### - Prepare the environment ###
    
### - Initialize Kubeflex ###

### - Install the Kubestellar core components ###

   
### - Create the Inventory & Transport Space ###

### - Create the Workload Description Spaces ###

## Create Workload Execution Clusters (WECs)

## Register WECs with Kubestellar core

## Create and Define workloads for deployment -->
