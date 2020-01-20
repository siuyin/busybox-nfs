# busybox-nfs is a demonstration of busybox with nfs persistent volumes.

1. First create the persistent volume (remember to edit pv.yaml for the correct server).
```
kubectl create -f pv.yaml
```

1. To deploy run:
```
kubectl apply -k base.  (kubectl version >= 1.14)
Or
kustomize build base/|kubectl create -f -
```
