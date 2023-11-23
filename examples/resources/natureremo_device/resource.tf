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

resource "natureremo_device" "mine" {
  name               = "Remo"
  temperature_offset = 0
  humidity_offset    = 0
}
