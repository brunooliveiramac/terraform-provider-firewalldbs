package firewalldbs

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dataprovider "terraform-provider-firewalldbs/firewalldbs/data_provider"
)

// Provider -
// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_CLIENT_SECRET", nil),
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_SUBSCRIPTION_ID", nil),
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_TENANT_ID", nil),
			},
			"agent_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AGENT_IP", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"firewalldbs_open":  resourceOpenFirewall(),
			"firewalldbs_close": resourceCloseFirewall(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, resource *schema.ResourceData) (interface{}, diag.Diagnostics) {

	clientId := resource.Get("client_id").(string)
	clientSecret := resource.Get("client_secret").(string)
	subscriptionId := resource.Get("subscription_id").(string)
	tenantId := resource.Get("tenant_id").(string)
	agentIp := resource.Get("agent_ip").(string)

	var diagnostics diag.Diagnostics

	credentials := dataprovider.Credential{
		GrantType:    "client_credentials",
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Resource:     "https://management.azure.com/",
		Tenant:       tenantId,
	}

	token, err := dataprovider.Login(&credentials)

	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create provider client connection the Cloud Provider",
			Detail:   err.Error(),
		})
		return nil, diagnostics
	}

	ip, err := dataprovider.GetIp(agentIp)

	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to obtain agent ip",
			Detail:   err.Error(),
		})
		return nil, diagnostics
	}

	connection := &dataprovider.Connection{
		Subscription: subscriptionId,
		Token:        token,
		AgentIP:      ip,
	}

	return connection, diagnostics
}
