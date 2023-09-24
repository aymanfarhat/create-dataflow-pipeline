module "docker_artifact_registry" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/artifact-registry"
  project_id = module.project.project_id
  name       = "${var.prefix}-df-containers"
  location   = var.region
  format = {
    docker = {
      immutable_tags = true
    }
  }
}
