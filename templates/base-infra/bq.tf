locals {
  table_schema = jsonencode([
    { name = "word", type = "STRING" },
    { name = "count", type = "INTEGER" }
  ])
}

module "bigquery-dataset" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/bigquery-dataset"
  project_id = var.project_id
  id         = "${var.prefix}_dataset"
  options = {
    delete_contents_on_destroy = true
  }
  tables = {
    messages = {
      friendly_name       = "WordCount"
      labels              = {}
      options             = null
      partitioning        = null
      schema              = local.table_schema
      deletion_protection = false
    }
  }
}

