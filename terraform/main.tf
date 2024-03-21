terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}
//Use the digitalocean Provider
provider "digitalocean" {
  token = var.token
}

//Use the digitalocean_kubernetes_cluster resource to create
//a Kubernetes cluster
resource "digitalocean_kubernetes_cluster" "foo" {
    version = var.k8s_version
    name = var.label
    region = var.region
    tags = var.tags

    node_pool {
        name       = "worker-pool"
        size       = "s-2vcpu-2gb"
        node_count = 3
    }
}

//Export this cluster's attributes
output "kubeconfig" {
  value = digitalocean_kubernetes_cluster.foo.kube_config
  sensitive = true
}

// output "api_endpoints" {
//  value = digitalocean_kubernetes_cluster.foo.api_endpoints
// }

output "status" {
  value = digitalocean_kubernetes_cluster.foo.status
}

output "id" {
  value = digitalocean_kubernetes_cluster.foo.id
}

// output "pool" {
//  value = digitalocean_kubernetes_cluster.foo.pool
// }