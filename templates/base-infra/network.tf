module "vpc" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/net-vpc"
  project_id = module.project.project_id
  name       = "${var.prefix}-df-vpc"
  subnets = [
    {
      ip_cidr_range = var.ip_ranges.dataflow
      name          = "subnet"
      region        = var.region
    }
  ]
}
