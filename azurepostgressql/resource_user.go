package azurepostgressql

import (
	"context"
	"fmt"
	"github.com/brunooliveiramac/azure-postgres-user-provider/azurepostgressql/data_provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDBUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive: 	 true,
			},
			"role": {
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	databaseCredentials := providerConfig.(data_provider.Credential)

	role := resource.Get("role").(string)
	username := resource.Get("username").(string)
	password := resource.Get("password").(string)

	roleToSearch := data_provider.Role{
		Name: role,
	}

	roleSelected, err := data_provider.SelectRole(&databaseCredentials, &roleToSearch)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find the role to associate to this new user",
			Detail:   msg,
		})

		return diagnostics
	}

	user := data_provider.User{
		Username: username,
		Password: password,
		Role:     *roleSelected,
	}

	_, err = data_provider.CreateUser(&databaseCredentials, &user)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Create User",
			Detail:   msg,
		})

		return diagnostics
	}

	resource.SetId(user.Username)

	resourceUserRead(ctx, resource, providerConfig)

	return diagnostics
}

func resourceUserRead(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	databaseCredentials := providerConfig.(data_provider.Credential)

	role := resource.Get("role").(string)
	username := resource.Get("username").(string)
	password := resource.Get("password").(string)

	roleToSearch := data_provider.Role{
		Name: role,
	}

	roleSelected, err := data_provider.SelectRole(&databaseCredentials, &roleToSearch)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to find the role to associate to this new user",
			Detail:   msg,
		})

		return diagnostics
	}

	user := data_provider.User{
		Username: username,
		Role:     *roleSelected,
	}

	userSelected, err := data_provider.SelectUser(&databaseCredentials, &user)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to select the user",
			Detail:   msg,
		})

		return diagnostics
	}

	if err := resource.Set("username", userSelected.Username); err != nil {
		return diag.FromErr(err)
	}
	if err := resource.Set("role", roleSelected.GetRoleName()); err != nil {
		return diag.FromErr(err)
	}
	if err := resource.Set("password", password); err != nil {
		return diag.FromErr(err)
	}

	return diagnostics
}

func resourceUserUpdate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	databaseCredentials := providerConfig.(data_provider.Credential)

	currentUsername := resource.Id()

	newUser := data_provider.User{
		Username:   currentUsername,
		NewUsername: resource.Get("username").(string),
		NewPassword: resource.Get("password").(string),
	}

	if resource.HasChange("username"){
		_, err := data_provider.AlterUsername(&databaseCredentials, &newUser)

		if err != nil {
			msg := fmt.Sprintf("%s", err)

			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to change username",
				Detail:   msg,
			})

			return diagnostics
		}

		resource.SetId(newUser.NewUsername)
	}

	if resource.HasChange("password"){
		_, err := data_provider.AlterPassword(&databaseCredentials, &newUser)

		if err != nil {
			msg := fmt.Sprintf("%s", err)

			diagnostics = append(diagnostics, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to change password",
				Detail:   msg,
			})

			return diagnostics
		}
	}

	return resourceUserRead(ctx, resource, providerConfig)
}

func resourceUserDelete(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	databaseCredentials := providerConfig.(data_provider.Credential)

	var diagnostics diag.Diagnostics

	username := resource.Get("username").(string)

	userToDrop := data_provider.User{
		Username:       username,
	}

	_, err := data_provider.DropUser(&databaseCredentials, &userToDrop)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable drop user",
			Detail:   msg,
		})
		return diagnostics
	}

	resource.SetId("")

	return diagnostics
}