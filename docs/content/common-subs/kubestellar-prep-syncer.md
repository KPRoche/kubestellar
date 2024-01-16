<!--kubestellar-prep-syncer-start-->
```shell hl_lines="3 7"
KUBECONFIG=~/.kube/config kubectl kubestellar prep-for-cluster --imw imw1 ks-edge-cluster1 \
  env=ks-edge-cluster1 \
  location-group=edge     #add ks-edge-cluster1 and ks-edge-cluster2 to the same group

KUBECONFIG=~/.kube/config kubectl kubestellar prep-for-cluster --imw imw1 ks-edge-cluster2 \
  env=ks-edge-cluster2 \
  location-group=edge     #add ks-edge-cluster1 and ks-edge-cluster2 to the same group
```
<!--kubestellar-prep-syncer-end-->