# Terraform Provider for Nature Remo

The Terrafrom Nature Remo provider is a plugin that allows [Terraform](https://www.terraform.io/) to manage resources on [Nature Remo Cloud API](https://developer.nature.global/).

This is the `natureremo` provider, containing generally available features.

`natureremo` provider is open to public in [Terraform Registry](https://registry.terraform.io/providers/aiwasaki126/natureremo/latest).


## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Building the Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the Provider

See provider [example](./docs/index.md).

## Developing the Provider

_This repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). The template repository built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) can be found at [terraform-provider-scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding)._

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
