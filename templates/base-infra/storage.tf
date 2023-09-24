module "pipeline_bucket" {
  source                      = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/gcs"
  project_id                  = var.project_id
  name                        = format("%s-%s", var.project_id, "df-pipeline")
  location                    = var.multi_regional
  uniform_bucket_level_access = true
  force_destroy               = true
  iam = {
    "roles/storage.objectAdmin" = [module.dataflow-sa.iam_email]
  }
  depends_on = [module.dataflow-sa]
}
