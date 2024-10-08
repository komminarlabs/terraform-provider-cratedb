---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "CrateDB Provider"
subcategory: ""
description: |-
  Use the CrateDB provider to deploy and manage resources supported by CrateDB. You must configure the provider with the proper credentials before you can use it.
---

# CrateDB Provider

Use the CrateDB provider to deploy and manage resources supported by CrateDB. You must configure the provider with the proper credentials before you can use it.


## Example Usage

```terraform
terraform {
  required_providers {
    cratedb = {
      source = "komminarlabs/cratedb"
    }
  }
}

provider "cratedb" {}
```

## Environment Variables

Credentials can be provided by using the `CRATEDB_API_KEY` and `CRATEDB_API_SECRET` and `CRATEDB_URL`.

### Example

```terraform
export CRATEDB_API_KEY="*******"
export CRATEDB_API_SECRET="*******"
export CRATEDB_URL="https://console.cratedb.cloud/"

provider "cratedb" {}

terraform plan
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String, Sensitive) The API key
- `api_secret` (String, Sensitive) The API secret
- `url` (String) The CrateDB Cloud URL
