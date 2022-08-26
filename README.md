# kubetest2-aks

Kubetest2-aks is a [kubetest2](https://github.com/kubernetes-sigs/kubetest2) deployer for Azure
Kubernetes Service clusters.

## Installation
To install core and all deployers and testers:
`GO111MODULE=on go get sigs.k8s.io/kubetest2/...@latest`

To install the AKS deployer:
`GO111MODULE=on go get github.com/juan-lee/kubetest2-aks@latest`

## Usage
An example run of the Ginkgo conformance suite on AKS looks as follows:
```
# authenticate
az login

# create resource group
az group create --location westus --name test-clusters

# run conformance
kubetest2 aks -v 2 \
  --resource-group test-clusters \
  --template ./templates/cluster.bicep \
  --up \
  --down \
  --test=ginkgo \
  -- \
  --focus-regex='\[Conformance\]'
```
