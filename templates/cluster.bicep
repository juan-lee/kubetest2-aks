param clusterName string = 'test'
param location string = resourceGroup().location

resource aksCluster 'Microsoft.ContainerService/managedClusters@2022-06-01' = {
  name: clusterName
  location: location
  sku: {
    name: 'Basic'
    tier: 'Paid'
  }
  identity: {
    type: 'SystemAssigned'
  }
  properties: {
    dnsPrefix: clusterName
    enableRBAC: true
    agentPoolProfiles: [
      {
        name: 'nodepool'
        count: 10
        vmSize: 'Standard_d8s_v3'
        mode: 'System'
      }
    ]
  }
}
