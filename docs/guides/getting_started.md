---
page_title: Getting Started
description: |-
  This guide explains how to manage Nature Remo settings using Nature Remo provider for Terraform.
---

# Getting Started

## Nature Remo

[Nature Remo](https://shop.nature.global/collections/nature-remo) is a smart remote control that allows you to control your home appliances using your smartphone or Apple Watch. It can be used with any home appliance equipped with an infrared remote control, regardless of make, model number, or year.



## Provider Setup

The provider needs to be cofigured with the access token.

```hcl
provider "natureremo" {
  access_token = var.access_token
}

variable "access_token" {
  type      = string
  sensitive = true
}
```

In `terraform.tfvars`, you can place the access token.

```hcl
access_token = "XXX"
```

## Import Your Profile

User profile cannot be created through Nature Remo API.
Because the profile only can be updated, we must import the user profile at first if you need to manage it.

### 1. Specify Your Profile ID

Using data source, specify your profile ID. Your ID will be printed on the cosole.

```hcl
data "natureremo_user" "me" {}

output "my_profile" {
  value = data.natureremo_user.me
}
```

### 2. Define Your Profile

Define your profile resource.

```hcl
resource "natureremo_user" "me" {
  nickname      = "REMO OWNER"
  country       = "JP"
  distance_unit = "metric"
  temp_unit     = "c"
}
```

### 3. Import Your Profile

Import your profile resource using your profile ID shown at step 2.

```hcl
terrafrom import natureremo_user.me YOUR_PROFILE_ID
```


## Import Your Nature Remo Device

Device cannot also be created. The same way as profile, you must import device resource at first.
