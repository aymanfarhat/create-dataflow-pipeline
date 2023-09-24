locals {}

module "project" {
  source         = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/project"
  name           = var.project_id
  project_create = false
  services = [
    "sql-component.googleapis.com",
    "vpcaccess.googleapis.com",
    "servicenetworking.googleapis.com",
    "compute.googleapis.com",
    "container.googleapis.com",
    "logging.googleapis.com",
    "monitoring.googleapis.com",
    "containerregistry.googleapis.com",
    "cloudbuild.googleapis.com",
    "storage.googleapis.com",
    "artifactregistry.googleapis.com",
    "sourcerepo.googleapis.com",
    "sqladmin.googleapis.com",
    "dataflow.googleapis.com",
  ]
}

