---
layout: "awx"
page_title: "AWX: awx_credential_type"
sidebar_current: "docs-awx-datasource-credential_type"
description: |-
  Use this data source to query Credential Type by ID.
---

# awx_credential_type

Use this data source to query Credential Type by ID.

## Example Usage

```hcl
data "awx_credential_type" "machine_ssh" {
  id = 1
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) ID of Credential Type

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - (Computed) Description of the credential type
* `injectors` - (Computed) Injector data for credential type
* `inputs` - (Computed) Input data for credential type
* `kind` - (Computed) Kind of credential type
* `name` - (Computed) Name of the credential type
