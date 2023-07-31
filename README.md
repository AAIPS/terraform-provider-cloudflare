# Cloudflare Terraform Provider

## Quickstarts

- [Getting started with Cloudflare and Terraform](https://developers.cloudflare.com/terraform/installing)
- [Developing the provider](contributing/development.md)

## Minimum requirements

- Terraform 1.2 or newer. We recommend running the [latest version](https://developer.hashicorp.com/terraform/downloads?product_intent=terraform) for optimal compatibility with the Cloudflare provider. Terraform versions older than 1.2 have known issues with newer features and internals.

## Documentation

Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs). [API documentation](https://api.cloudflare.com) and [Developer documentation](https://developers.cloudflare.com) is also available
for non-Terraform or service specific information.

## Migrating to Terraform from using the Dashboard

Do you have an existing Cloudflare account (or many!) that you'd like to transition
to be managed via Terraform? Check out [cf-terraforming](https://github.com/cloudflare/cf-terraforming)
which is a tool Cloudflare has built to help dump the existing resources and
import them into Terraform.

## Contributing

To contribute, please read the [contribution guidelines](contributing/README.md).

## Feedback

If you would like to provide feedback (not a bug or feature request) on the Cloudflare Terraform provider, you're welcome to via [this form](https://forms.gle/6ofUoRY2QmPMSqoR6).

## Development

For local testing let's setup dev overrides (see [docs](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install))

Create file `.terraformrc` in your home directory with content:

```
provider_installation {

  dev_overrides {
      "localhost/cloudflare/cloudflare" = "<PATH_TO_DIR_WITH_PROVIDER_EXECUTABLE>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

`PATH_TO_DIR_WITH_PROVIDER_EXECUTABLE` need to be replaced with a correct path on your system, where go installs binaries,
typically it's $GOBIN or $GOPATH/bin.

`localhost/cloudflare/cloudflare` - full name of our provider, we are setting it in main.go:27

To be able build and install provider locally, you can run `make install`, this will create new binary for the provider 
and override existing executable `terraform-provider-cloudflare` in $GOBIN path.

## Example: provider

Now let's open first example, navigate to `examples/provider` and run:

```bash
terraform plan

# or

terraform apply
```

You will see similar output:

```
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - localhost/cloudflare/cloudflare in <PATH_TO_DIR_WITH_PROVIDER_EXECUTABLE>
│ 
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become incompatible with published releases.
╵
╷
│ Error: Invalid Attribute Value Match
│ 
│   with provider["localhost/cloudflare/cloudflare"],
│   on provider.tf line 10, in provider "cloudflare":
│   10:   api_token = "<API_TOKEN>"
│ 
│ Attribute api_token API tokens must be 40 characters long and only contain characters a-z, A-Z, 0-9, hyphens and underscores, got: <API_TOKEN>
╵
```

This means our dev overrides works, and we can use our local provider executable.
`Error: Invalid Attribute Value Match` means validation rule defined for `api_token` detected issues in the input data.

If you set the correct API_TOKEN, you can see the error, which means `cloudflare_ruleset` resource is not defined in the provider:

```
│ Error: Invalid resource type
│ 
│   on provider.tf line 14, in resource "cloudflare_ruleset" "example":
│   14: resource "cloudflare_ruleset" "example" {}
│ 
│ The provider localhost/cloudflare/cloudflare does not support resource type "cloudflare_ruleset".
╵
```