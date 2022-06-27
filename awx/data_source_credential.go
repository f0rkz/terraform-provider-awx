/*
Use this data source to query Credential by ID.

Example Usage

```hcl
data "awx_credential" "example" {
  id = 1
}
```

*/
package awx

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/mrcrilly/goawx/client"
)

func dataSourceCredentialByID() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCredentialByIDRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Description: "ID of Credential",
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Description: "(Computed) Username from inputs username field",
				Computed: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Description: "(Computed) Kind of credential",
				Computed: true,
			},
		},
	}
}

func dataSourceCredentialByIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*awx.AWX)
	id := d.Get("id").(int)
	cred, err := client.CredentialsService.GetCredentialsByID(id, map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to fetch credential",
			Detail:   "The given credential ID is invalid or malformed",
		})
	}

	d.Set("username", cred.Inputs["username"])
	d.Set("kind", cred.Kind)
	d.Set("tower_id", id)
	d.SetId(strconv.Itoa(id))

	return diags
}
