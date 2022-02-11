package azurepostgressql

import (
	"context"
	"github.com/brunooliveiramac/azure-postgres-user-provider/azurepostgressql/data_provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDBRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRoleCreate,
		ReadContext:   resourceRoleRead,
		UpdateContext: resourceRoleUpdate,
		DeleteContext: resourceRoleDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"privileges": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The list of privileges to grant",
			},
		},
	}
}

func resourceRoleCreate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	databaseCredentials := providerConfig.(*data_provider.Credential)

	role := resource.Get("role").(string)

	var privileges []string

	for _, privilege := range resource.Get("privileges").(*schema.Set).List() {
		privileges = append(privileges, privilege.(string))
	}

	roleWithPrivilege := data_provider.Role{
		Name:       role,
		Privileges: privileges,
	}

	data_provider.CreateRole(databaseCredentials, &roleWithPrivilege)

	return diagnostics
}

func resourceRoleRead(ctx context.Context, resource*schema.ResourceData, m interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	return diagnostics
}

func resourceRoleUpdate(ctx context.Context, resource *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceRoleRead(ctx, resource, m)
}

func resourceRoleDelete(ctx context.Context, resource*schema.ResourceData, m interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	return diagnostics
}