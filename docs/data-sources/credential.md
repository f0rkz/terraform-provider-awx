---
layout: "awx"
page_title: "AWX: awx_credential"
sidebar_current: "docs-awx-datasource-credential"
description: |-
  Use this data source to query Credential by ID.
---

# awx_credential

Use this data source to query Credential by ID.

## Example Usage

```hcl
data "awx_credential" "example" {
	id = 1
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) ID of Credential

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `kind` - (Computed) Kind of credential
* `username` - (Computed) Username from inputs username field
