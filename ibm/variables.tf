variable "ibmcloud_api_key" {
  description = "IBM Cloud API Key. (required)"
}

variable "k8s_version" {
  description = "The Kubernetes version to use for this cluster. (required)"
  default     = "4.7_openshift"
}

variable "resource_group" {
  description = "Resource group name where the cluster will be created. (required)"
  default     = "Default"
}

variable "name" {
  description = "The unique name to assign to this Kubernetes cluster. (required)"
  default     = "ikc"
}

//variable "vpc_name" {
//  description = "The name of the VPC where the cluster will be located. (required)"
//}

//variable "subnet_zone" {
//  description = "The zone where the subnet will be created. (required)"
//}

variable "subnet_cidr" {
  description = "The CIDR block for the subnet. (required)"
  default     = "10.240.0.0/24"
}

variable "machine_type" {
  description = "The machine type for the Kubernetes workers. (required)"
  default     = "cx2.2x4"
}

variable "worker_count" {
  description = "The number of worker nodes for the cluster. (required)"
  default     = 2
}

variable "worker_zones" {
  description = "The zones where the worker nodes will be created. (required)"
  default     = ["us-south-1"]
}

variable "default_worker_pool_size" {
  description = "The default size of the worker pool. (required)"
  default     = 2
}