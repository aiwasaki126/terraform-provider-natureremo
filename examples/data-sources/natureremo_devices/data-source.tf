terraform {
  required_providers {
    natureremo = {
      source = "registry.terraform.io/hashicorp/natureremo"
    }
  }
}

provider "natureremo" {
  access_token = var.access_token
}

variable "access_token" {
  type      = string
  sensitive = true
}

data "natureremo_devices" "mine" {}

output "my_devices" {
  value = data.natureremo_devices.mine
}
