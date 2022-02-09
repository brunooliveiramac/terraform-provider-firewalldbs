package azurepostgressql

import (
	"context"
	dataprovider "github.com/brunooliveiramac/azure-postgres-user-provider/azurepostgressql/data_provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive: 	  false,
			},
			"database": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive: 	  true,
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive: 	 true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azurepostgressql_user": resourceDBUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, resource *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := resource.Get("host").(string)
	port := resource.Get("port").(int)
	username := resource.Get("username").(string)
	password := resource.Get("password").(string)
	database := resource.Get("database").(string)

	var diagnostics diag.Diagnostics

	credentials := dataprovider.Credential{
		Host:     host,
		Database: database,
		Port:     port,
		Username: username,
		Password: password,
	}

	connection, err := dataprovider.DBClient(&credentials)

	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create provider client connection",
			Detail:   "Unable to auth user on server",
		})
		return nil, diagnostics
	}

	diagnostics = append(diagnostics, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Warning Message Summary",
		Detail:   "This is the detailed warning message from providerConfigure",
	})

	return connection, diagnostics
}


