package firewalldbs

import (
	"context"
	"fmt"
	"github.com/brunooliveiramac/azure-postgres-user-provider/firewalldbs/data_provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloseFirewall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloseFirewallCreate,
		ReadContext:   resourceCloseFirewallRead,
		UpdateContext: resourceCloseFirewallUpdate,
		DeleteContext: resourceCloseFirewallDelete,
		Schema: map[string]*schema.Schema{
			"server_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"agent_ip": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AGENT_IP", nil),
			},
		},
	}
}

func resourceCloseFirewallCreate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	connection := providerConfig.(*data_provider.Connection)

	serverName := resource.Get("server_name").(string)
	resourceGroup := resource.Get("resource_group_name").(string)

	firewallRule := data_provider.ServerFirewallIpRule{
		IP:            connection.AgentIP,
		ServerName:    serverName,
		ResourceGroup: resourceGroup,
		Subscription:  connection.Subscription,
	}

	err := data_provider.DeleteAgentIp(&firewallRule, connection.Token)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Add Agent IP",
			Detail:   msg,
		})

		return diagnostics
	}

	resource.SetId("AgentIP")

	resourceCloseFirewallRead(ctx, resource, providerConfig)

	return diagnostics
}

func resourceCloseFirewallRead(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	connection := providerConfig.(*data_provider.Connection)

	serverName := resource.Get("server_name").(string)
	resourceGroup := resource.Get("resource_group_name").(string)

	firewallRule := data_provider.ServerFirewallIpRule{
		IP:            connection.AgentIP,
		ServerName:    serverName,
		ResourceGroup: resourceGroup,
		Subscription:  connection.Subscription,
	}

	err := data_provider.DeleteAgentIp(&firewallRule, connection.Token)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable To Delete Rule",
			Detail:   msg,
		})
		return diagnostics
	}

	if err := resource.Set("agent_ip", connection.AgentIP); err != nil {
		return diag.FromErr(err)
	}

	resource.SetId("AgentIP")

	return diagnostics
}

func resourceCloseFirewallUpdate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics


	resource.SetId("AgentIP")

	connection := providerConfig.(*data_provider.Connection)


	if err := resource.Set("agent_ip", connection.AgentIP); err != nil {
		return diag.FromErr(err)
	}

	return diagnostics
}

func resourceCloseFirewallDelete(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	resource.SetId("")

	diagnostics = append(diagnostics, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Removing AgentName from tfState",
		Detail:   "Info: The delete does not actually remove anything, To do it use the close resource.",
	})
	return diagnostics

	return diagnostics
}
