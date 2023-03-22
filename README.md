# epilotpermissions

This is a Pre-Alpha release of the epilotpermissions Terraform Provider.

This provider is not intended for production use at this time.

## Installing the Provider

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```
terraform {
  required_providers {
    epilotpermissions = {
      source  = "epilot/epilotpermissions"
      version = "0.0.1"
    }
  }
}

provider "epilotpermissions" {
  # Configuration options
}
```

## Testing the Provider

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```