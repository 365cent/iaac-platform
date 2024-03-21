terraform {
  required_providers {
<<<<<<< HEAD
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
=======
    linode = {
      source = "linode/linode"
      version = "2.7.1"
    }
  }
}
//Use the Linode Provider
provider "linode" {
  token = var.token
}

//Use the linode_lke_cluster resource to create
//a Kubernetes cluster
resource "linode_lke_cluster" "foobar" {
    k8s_version = var.k8s_version
    label = var.label
    region = var.region
    tags = var.tags

    dynamic "pool" {
        for_each = var.pools
        content {
            type  = pool.value["type"]
            count = pool.value["count"]
        }
>>>>>>> refs/remotes/origin/main
    }
}

//Export this cluster's attributes
output "kubeconfig" {
<<<<<<< HEAD
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
=======
  value = linode_lke_cluster.foobar.kubeconfig
  sensitive = true
}

output "api_endpoints" {
  value = linode_lke_cluster.foobar.api_endpoints
}

output "status" {
  value = linode_lke_cluster.foobar.status
}

output "id" {
  value = linode_lke_cluster.foobar.id
}

output "pool" {
  value = linode_lke_cluster.foobar.pool
}
>>>>>>> refs/remotes/origin/main
