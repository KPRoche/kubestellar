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
  - Prep the mailbox controller for and create the yaml files to build syncers for the WECs
  - Run the -syncer.yaml files on the WEC host(s) to create the syncer for each WEC
  - Set up the workload description on the core
  - Set up the  

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
4. Deploy kubestellar core into a kubernetes cluster
    - Create a cluster with appropriate ingress (link to add)
    - use **kubectl kubestellar deploy** command (with necessary flags and options) will deploy all the core components to either a kind cluster or (if flag is set) an openshift cluster.
    - fetch the external kubeconfig and internal kubeconfig files for the kubestellar server
    - if creating WEC on a different host, modify port control for the server as needed
5. Create the inventory management workspace on your kubestellar core
    - **use the correct kubeconfig file**
       - that will be the file fetched via the _kubectl kubestellar get-external-kubeconfig_ command
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
## Install tools on the WEC system
  - The WEC systems must support creation of Kubernetes clusters
  - The WECs will need TCP access to the kubestellar core host, and to public container images on the registry at quay.io
  
## Create or get access info for the WEC(s)
  - In my experiment, I created the florin and guilder clusters with Kind as described in the extended example
  - ?? You will need to copy the _core_external.kubeconfig_ file created from the Kubestellar core system over to the WEC host ??
  - The WECs will need TCP access to the kubestellar core host, and to public container images on the registry at quay.io

## Set up the objects on the core to correspond to the WECs
**All these steps take place on the system hosting Kubestellar core!**

**_From this point on notes are for a kubestellar core running in a kubernetes cluster_** 
  - Follow the procedure described in the Extended Example 
  - For safety's sake, note and save the path to your standard (not Kubestellar server) kubeconfig file
  - export KUBECONFIG=(_path to external_kubeconfig file_ )
(easy way to do that is to change to the folder holding it and `export KUBECONFIG=$(pwd)/config_filename`)
### Create the location and SyncTarget objects
  - execute commands to create location and SyncTarget objects in ks-core for each WEC:
    - `kubectl ws root:imw-1` [specifies the definition space used for next commands]
    - `kubectl kubestellar ensure location florin  loc-name=florin  env=prod` This creates the objects for a simple 1 workload WEC
    - `kubectl kubestellar ensure location guilder loc-name=guilder env=prod extended=si` This creates the objects for a WEC that will receive more than one workload (guilder in the example was configured for two)
    - `kubectl describe location.edge.kubestellar.io florin` will fetch the object description for the location "florin"
    - `kubectl describe location.edge.kubestellar.io guilder` will do the same for guilder
    - or, all together:
      ```
      kubectl ws root:imw-1
      kubectl kubestellar ensure location florin  loc-name=florin  env=prod
      kubectl kubestellar ensure location guilder loc-name=guilder env=prod extended=si
      echo "describe the location objects created"
      kubectl describe location.edge.kubestellar.io florin
      kubectl describe location.edge.kubestellar.io guilder
      ```
## Prep the mailbox controller for and Create the yaml files to build syncers for the WECs
**still on the kubestellar core hosting system, with $KUBECONFIG pointing at the correct configuration**
  - For each WEC, this command will create the id and authorization in its mailbox workspace and output a 
   **cluster-syncer.yaml** file to run **ON THE WEC host**:

    _kubectl kubestellar prep-for-syncer --imw root:imw-1 cluster_

  - for example, for my example 1 clusters guilder:
    `kubectl kubestellar prep-for-syncer --imw root:imw-1 guilder`
    will set up the identity and authorization in its mailbox space and write the file _guilder-syncer.yaml_ in the current directory.

## Copy the -syncer.yaml file(s) and use them to create syncers on the WEC(s)
  - Copy each _xxx-syncer.yaml_ (where _xxx_ is the cluster name) file to the host machine on which that WEC is running
  - in the same kubeconfig context in which the WEC is running, apply the yaml file
  
     _kubectl --context [_xxx_] apply -f xxx-syncer.yaml_

     to create the KubeStellar syncer in the cluster _xxx_
  - for the example clusters, it looks like
   
      **kubectl --context kind-guilder apply -f guilder-syncer.yaml**   and    
      **kubectl --context kind-florin apply -f florin-syncer.yaml**

## Set up the workload descriptions in the core
**Working on the _core_ host, using the external_kubeconfig file**
  - make sure you are working in the workload management space
    ```
    kubectl ws root
    kubectl kubestellar ensure wmw wmw-c
    ```
  - use the files* common.yaml and edgeplacement.yaml to create the entries for the first workload
    ```
    kubectl apply -f common.yaml
    kubectl apply -f edgeplacement.yaml
    ```
  - use the files* special.yaml and edgeplacement-s.yaml to create entries for the special second workload
    ```
    kubectl apply -f special.yaml
    kubectl apply -f edgeplacement-s.yaml
    ```

    * files created from the command texts on the extended example page


 
