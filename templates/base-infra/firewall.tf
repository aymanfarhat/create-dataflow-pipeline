module "firewall" {
  source     = "github.com/GoogleCloudPlatform/cloud-foundation-fabric.git//modules/net-vpc-firewall"
  project_id = module.project.project_id
  network    = module.vpc.name

  egress_rules = {
    allow-egress-df = {
      description = "Allow outgoing tcp requests to 12345-12346"
      deny        = false
      rules = [
        {
          protocol = "tcp"
          ports    = ["12345", "12346"]
        },
      ]
    }

    allow-ingress-df = {
      description = "Allow incoming tcp requests to 12345-12346"
      deny        = false
      rules = [
        {
          protocol = "tcp"
          ports    = ["12345", "12346"]
        },
      ]
    }
  }
}
