variable "subnetwork_name" {
  type    = "string"
  default = "default"
}

variable "zone" {
  type    = "string"
  default = "us-east1-b"
}

variable "machine_type" {
  type    = "string"
  default = "n1-standard-2"
}

variable "disk_type" {
  type        = "string"
  description = "The type of volume for the root block device."
  default     = "pd-standard"
}

variable "disk_size" {
  type        = "string"
  default     = "60"
  description = "The size of the volume in gigabytes for the root block device."
}

variable "public_ssh_key" {
  default = ""
}

variable "gcp_service_account" {
  default = ""
}

variable "instance_count" {
  default = 1

}
