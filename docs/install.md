# Install Rio

Rio can be installed from both CLI or kubernetes manifest.

1. Download the latest release. 

```bash
$ curl -sfL https://get.rio.io | sh - 
```

Note: by default it will only download the latest release from github. To test a pre-release or alpha build, run.

```bash
$ curl -sfL https://get.rio.io | INSTALL_RIO_VERSION=${version} sh - 
```

2. Prepare a Kubernetes cluster, see [Clusters and Providers](#clusters-and-providers). Setup KUBECONFIG environment variable to point to your kubernetes cluster.

3. Run 

```bash
$ rio install
```

Note: to install rio from kubernetes manifest, run `rio install --yaml`. It will print out kubernetes manifest instead of 
installing rio directly, so that you can apply the manifest later.

## Options

See the [CLI install reference docs](cli-reference.md#install) for complete list of options.

## Clusters and Providers

A 1.15 or higher version of kubernetes is recommended.

See the kubernetes [getting started guide](https://kubernetes.io/docs/setup/) to help choose a cluster type.

**Internal vs external IP address**

Some cloud providers will use an internal IP address by default and Rio will use that for its cluster domain.
You can override this during cluster setup by configuring external IP's, or by using a cloud-provider to setup a load balancer which will have an external IP.

For an example of installing a cluster with a cloud provider on Linode using Rancher, see [bullet #8 on this doc](https://www.linode.com/docs/kubernetes/how-to-deploy-kubernetes-on-linode-with-rancher-2-x/#provision-a-cluster).

**k3s**

When installing on k3s use the `--no-deploy traefik` flag.

**EKS**

Ensure you are running enough nodes with proper instance types to allow for the rio and kubernetes systems to run at least 45 pods.
See the docs to help determine proper sizes:

* Guide: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI
* Total limits: https://github.com/awslabs/amazon-eks-ami/blob/master/files/eni-max-pods.txt

For example: you have a workload that requires 15 pods and you want to run 2 nodes.
15+45 is a 60 pod minimum requirement, so running 2 t3.large nodes (2*35=70) would be just enough.