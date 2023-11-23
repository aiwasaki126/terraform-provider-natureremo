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

resource "natureremo_user" "me" {
  nickname      = "REMO OWNER"
  country       = "JP"
  distance_unit = "metric"
  temp_unit     = "c"
}
