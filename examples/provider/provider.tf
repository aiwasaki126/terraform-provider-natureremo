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

data "natureremo_user" "me" {}

output "me" {
  value = data.natureremo_user.me
}

variable "access_token" {
  type      = string
  sensitive = true
}
