variable "project_id" {
  description = "Project id, references existing project if `project_create` is null."
  type        = string
}

variable "dataset" {
  type        = string
  description = "Dataset to write results into."
}

variable "prefix" {
  description = "Prefix used for resource names."
  type        = string
  validation {
    condition     = var.prefix != ""
    error_message = "Prefix cannot be empty."
  }
}

variable "region" {
  type        = string
  description = "Region for the created resources."
  default     = "europe-west1"
}

variable "multi_regional" {
  type        = string
  description = "Continent for multi-regional resources such as buckets"
  default     = "EU"
}


variable "ip_ranges" {
  description = "CIDR blocks: Dataflow, CloudSQL etc..."
  type = object({
    dataflow = string
  })
  default = {
    dataflow = "10.0.0.0/24"
  }
}
