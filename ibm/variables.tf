variable "ibmcloud_api_key" {
  description = "IBM Cloud API Key. (required)"
}

variable "k8s_version" {
  description = "The Kubernetes version to use for this cluster. (required)"
  default     = "4.7_openshift"
}

variable "resource_group" {
  description = "Resource group name where the cluster will be created. (required)"
}

variable "cluster_name" {
  description = "The unique name to assign to this Kubernetes cluster. (required)"
}

variable "vpc_name" {
  description = "The name of the VPC where the cluster will be located. (required)"
}

variable "subnet_zone" {
  description = "The zone where the subnet will be created. (required)"
}

variable "subnet_cidr" {
  description = "The CIDR block for the subnet. (required)"
  default     = "10.240.0.0/24"
}

variable "machine_type" {
  description = "The machine type for the Kubernetes workers. (required)"
  default     = "bx2.4x16"
}

variable "worker_count" {
  description = "The number of worker nodes for the cluster. (required)"
  default     = 2
}
