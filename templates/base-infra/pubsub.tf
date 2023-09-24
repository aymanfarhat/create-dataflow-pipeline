module "pubsub_resources" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/pubsub"
  project_id = var.project_id
  name       = "${var.prefix}-topic"
  subscriptions = {
    "${var.prefix}-sub" = null
  }
}
