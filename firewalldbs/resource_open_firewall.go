package firewalldbs

import (
	"context"
	"fmt"
	"github.com/dchest/uniuri"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-firewalldbs/firewalldbs/data_provider"
)

func resourceOpenFirewall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOpenFirewallCreate,
		ReadContext:   resourceOpenFirewallRead,
		UpdateContext: resourceOpenFirewallUpdate,
		DeleteContext: resourceOpenFirewallDelete,
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

func resourceOpenFirewallCreate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

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

	err := data_provider.AddAgentIp(&firewallRule, connection.Token)

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

	resourceOpenFirewallRead(ctx, resource, providerConfig)

	return diagnostics
}

func resourceOpenFirewallRead(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	randomID := uniuri.New()

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

	err := data_provider.AddAgentIp(&firewallRule, connection.Token)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Add Agent IP",
			Detail:   msg,
		})

		return diagnostics
	}

	// On Read we have to change the current state before apply, so it will say that changes were made
	// outside terraform
	// Read expects the infra to be the same as the tfState

	ipName := fmt.Sprintf("%s_%s", connection.AgentIP, randomID)

	if err := resource.Set("agent_ip", ipName); err != nil {
		return diag.FromErr(err)
	}

	resource.SetId("AgentIP")

	return diagnostics
}

func resourceOpenFirewallUpdate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics


	resource.SetId("AgentIP")

	connection := providerConfig.(*data_provider.Connection)


	if err := resource.Set("agent_ip", connection.AgentIP); err != nil {
		return diag.FromErr(err)
	}

	return diagnostics
}

func resourceOpenFirewallDelete(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	resource.SetId("")

	diagnostics = append(diagnostics, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Removing AgentName from tfState",
		Detail:   "Info: The deletion does not actually remove anything. To do it use the close resource.",
	})
	return diagnostics

	return diagnostics
}
