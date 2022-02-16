package azurepostgressql

import (
	"context"
	"fmt"
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
			"schema": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"database": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"tables": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The list of objects to grant",
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

	databaseCredentials := providerConfig.(data_provider.Credential)

	role := resource.Get("name").(string)
	schma := resource.Get("schema").(string)

	var privileges []string
	var tables []string

	for _, privilege := range resource.Get("privileges").(*schema.Set).List() {
		privileges = append(privileges, privilege.(string))
	}

	for _, table := range resource.Get("tables").(*schema.Set).List() {
		tables = append(tables, table.(string))
	}

	roleWithPrivilege := data_provider.Role{
		Name:       role,
		Privileges: privileges,
		Schema: schma,
		Tables: tables,
	}

	_, err := data_provider.CreateRole(&databaseCredentials, &roleWithPrivilege)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Create Role",
			Detail:   msg,
		})

		return diagnostics
	}

	if len(roleWithPrivilege.Tables) == 1 && roleWithPrivilege.Tables[0] == "ALL" {
		_, err := data_provider.GrantPrivilegesOnAllTables(&databaseCredentials, &roleWithPrivilege)
		if err != nil {
			msg := fmt.Sprintf("%s", err)

			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable Grant Privileges for All Tables",
				Detail:   msg,
			})
			return diagnostics
		}
	} else if len(roleWithPrivilege.Tables) > 0 {
		_, err := data_provider.GrantPrivileges(&databaseCredentials, &roleWithPrivilege)
		if err != nil {
			msg := fmt.Sprintf("%s", err)

			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable Grant Privileges",
				Detail:   msg,
			})
			return diagnostics
		}
	}

	resource.SetId(roleWithPrivilege.GetRoleName())

	resourceRoleRead(ctx, resource, providerConfig)

	return diagnostics
}

func resourceRoleRead(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	// HERE TODO
	databaseCredentials := providerConfig.(data_provider.Credential)

	role := resource.Get("name").(string)

	var privileges []string
	var tables []string

	for _, privilege := range resource.Get("privileges").(*schema.Set).List() {
		privileges = append(privileges, privilege.(string))
	}

	for _, table := range resource.Get("tables").(*schema.Set).List() {
		tables = append(tables, table.(string))
	}

	roleWithPrivilege := data_provider.Role{
		Name:       role,
		Privileges: privileges,
	}

	selectedRole, err := data_provider.SelectRole(&databaseCredentials, &roleWithPrivilege)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable Select Role",
			Detail:   msg,
		})
		return diagnostics
	}

	if err := resource.Set("schema", selectedRole.GetRoleSchema()); err != nil {
		return diag.FromErr(err)
	}
	if err := resource.Set("name", selectedRole.GetRoleName()); err != nil {
		return diag.FromErr(err)
	}
	if len(selectedRole.Tables) == 1 && selectedRole.Tables[0] == "ALL" {
		if err := resource.Set("tables", selectedRole.Tables); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := resource.Set("privileges", selectedRole.Privileges); err != nil {
		return diag.FromErr(err)
	}

	resource.SetId(selectedRole.GetRoleName())

	return diagnostics
}

func resourceRoleUpdate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics
	role := resource.Get("name").(string)
	scm := resource.Get("schema").(string)

	var privileges []string
	var tables []string

	for _, privilege := range resource.Get("privileges").(*schema.Set).List() {
		privileges = append(privileges, privilege.(string))
	}

	for _, table := range resource.Get("tables").(*schema.Set).List() {
		tables = append(tables, table.(string))
	}

	roleWithPrivilegeAndTables := data_provider.Role{
		Name:       role,
		Privileges: privileges,
		Schema: scm,
	}

	databaseCredentials := providerConfig.(data_provider.Credential)

	_, err := data_provider.RevokeAll(&databaseCredentials, &roleWithPrivilegeAndTables)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable Update Role",
			Detail:   msg,
		})
		return diagnostics
	}

	if len(roleWithPrivilegeAndTables.Tables) > 0 {
		_, err := data_provider.GrantPrivileges(&databaseCredentials, &roleWithPrivilegeAndTables)
		if err != nil {
			msg := fmt.Sprintf("%s", err)

			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable Grant Privileges",
				Detail:   msg,
			})
			return diagnostics
		}
	} else {
		_, err := data_provider.GrantPrivilegesOnAllTables(&databaseCredentials, &roleWithPrivilegeAndTables)
		if err != nil {
			msg := fmt.Sprintf("%s", err)

			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable Grant Privileges for All Tables",
				Detail:   msg,
			})
			return diagnostics
		}
	}

	return resourceRoleRead(ctx, resource, providerConfig)
}

func resourceRoleDelete(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	databaseCredentials := providerConfig.(data_provider.Credential)

	var diagnostics diag.Diagnostics

	role := resource.Get("name").(string)
	scm := resource.Get("schema").(string)

	roleToDelete := data_provider.Role{
		Name:       role,
		Schema: scm,
	}

	_, err := data_provider.RevokeAll(&databaseCredentials, &roleToDelete)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to revoke privileges to delete Role",
			Detail:   msg,
		})
		return diagnostics
	}

	_, err = data_provider.DropRole(&databaseCredentials, &roleToDelete)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable drop Role",
			Detail:   msg,
		})
		return diagnostics
	}

	resource.SetId("")

	return diagnostics
}