variable "token" {
<<<<<<< HEAD
  description = "Your Digitalocean API Personal Access Token. (required)"
=======
  description = "Your Linode API Personal Access Token. (required)"
>>>>>>> refs/remotes/origin/main
}

variable "k8s_version" {
  description = "The Kubernetes version to use for this cluster. (required)"
<<<<<<< HEAD
  default = "1.29.1-do.0"
=======
  default = "1.28"
>>>>>>> refs/remotes/origin/main
}

variable "label" {
  description = "The unique label to assign to this cluster. (required)"
  default = "default-lke-cluster"
}

variable "region" {
  description = "The region where your cluster will be located. (required)"
<<<<<<< HEAD
  default = "nyc1"
=======
  default = "us-east"
>>>>>>> refs/remotes/origin/main
}

variable "tags" {
  description = "Tags to apply to your cluster for organizational purposes. (optional)"
  type = list(string)
  default = ["testing"]
}

variable "pools" {
  description = "The Node Pool specifications for the Kubernetes cluster. (required)"
  type = list(object({
    type = string
    count = number
  }))
  default = [
    {
      type = "g6-standard-1"
      count = 3
    }
  ]
}