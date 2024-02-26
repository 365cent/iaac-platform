variable "token" {
  description = "Your Linode API Personal Access Token. (required)"
  default = "0a5c8ce5613ae1be44952c0935f8ad19198d05c3f5d8564178bc1c4a634987f4"
}

variable "k8s_version" {
  description = "The Kubernetes version to use for this cluster. (required)"
  default = "1.28"
}

variable "label" {
  description = "The unique label to assign to this cluster. (required)"
  default = "default-lke-cluster"
}

variable "region" {
  description = "The region where your cluster will be located. (required)"
  default = "us-east"
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