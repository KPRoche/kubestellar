# Checklists for installing and setting up Kubestellar

## Basics
A Kubestellar system consists of two principal components, Kubestellar Core, which runs on a management system, and Kubestellar Syncer, which runs on each system to which a workload is being managed by the core, referred to as a _Workload Execution Cluster_ or WEC. (A _system_ can be a physical or virtual machine)

- **Kubestellar Core**  manages workload placement, and has several subcomponents
  - Kubestellar server which oversees its operations
  - workload and inventory databases to define (respectively) applications and the locations to place (deploy) them
  - an engine of some kind to do the placement (MCCM = Multi-Cluster Configuration Management)
  - another engine to monitor and report the status of the placed workloads
  - is designed to be technology agnostic for what placement/monitor engines are used. Currently is using kcp & a mailbox system, with one mailbox per WEC

  Kubestellar core can run directly on the host system ("bare") or in a kubernetes cluster on the host system

- **Kubestellar Syncer**
  - communicates with Kubestellar core
  - runs in a cluster (the _Workload Execution Cluster_) on the managed location
  - does actual deployment of the workload as mandated by Kubestellar core
  - reports status of the deployment back to core
  - accepts updates/reconfigurations/redeployments of workload from core
  - **does not** handle the data flow of the workload application itself, that is done by the application

## Order of setup
  - Install or gain access to the kubestellar core system
  - Install tools on the WEC system
  - Create or get access information for the WEC(s)
  - Create objects on the core to represent the WECs
  - 

## Install or gain access to the kubestellar core system

To use Kubestellar, you need a system running the kubestellar core and a system on which to run the cluster with the syncer.

_**Note: These can be the same physical system**_
### Installation
If you need to deploy the kubestellar core components
#### On system hosting kubestellar core
1. install prerequisites as listed in quickstart page
2. install release 0.11 of kcp for your language/architecture and add to $PATH 
    - download from https://github.com/kcp-dev/kcp/releases/tag/v0.11.0
    - expand files -- use tar CLI not GUI to ensure symlinks are created
    - add directories to $PATH
3. install stable release of kubestellar
    - download from https://github.com/kubestellar/kubestellar/releases
    - expand file
    - add directory to path
4. now you can follow instructions for running bare or using kubectl kubestellar deploy into a cluster
    - Running Bare
      - **NOTE:** if running bare, starting kcp will create the .kcp folder holding admin.kubeconfig as a subfolder of whatever folder you are in when you issue the command
    - Running in a cluster
      - You need to create a cluster with appropriate ingress (link to add)
       - use **kubectl kubestellar deploy** command (with necessary flags and options) will deploy all the core components to either a kind cluster or (if flag is set) an openshift cluster.
       - fetch the external kubeconfig and internal kubeconfig files for the kubestellar server
       - if creating WEC on a different host, modify port control for the server
5. Create the inventory management workspace on your kubestellar core
    - **use the correct kubeconfig file**
       - if running in a cluster, that will be the file fetched via the _kubectl kubestellar get-external-kubeconfig_ command
       - if running bare, make sure the context is set correctly
    - issue the commands 
      ```
      kubectl ws root
      kubectl ws create imw-1
      ```
    - you can check whether it worked with the command `kubectl ws tree` . It should return
      ```
      └── root
         ├── compute
         ├── espw
         └── imw-1
      ```
### Access
1. Install (WHAT COMPONENTS DO WE NEED ON THE USERS SYSTEM?) 
For now let us say do steps 1-3 in the Installation checklist
3. obtain the external kubeconfig file for the KubeStellar core system, eg _core_external.kubeconfig_
4. Test the connection with you can check whether it worked with the command `kubectl ws tree --kubeconfig=core_external.kubeconfig`. 
It should return something like
      ```
      └── root
         ├── compute
         ├── espw
         └── imw-1
      ```

 
