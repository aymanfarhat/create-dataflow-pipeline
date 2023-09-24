module "dataflow-sa" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/iam-service-account"
  project_id = var.project_id
  name       = "${var.prefix}-dataflow-sa"
  iam_project_roles = {
    "${var.project_id}" = [
      "roles/artifactregistry.reader",
      "roles/bigquery.dataEditor",
      "roles/bigquery.jobUser",
      "roles/cloudprofiler.agent",
      "roles/compute.networkUser",
      "roles/dataflow.worker",
      "roles/dataflow.admin",
      "roles/storage.objectAdmin",
    ]
  }
}
