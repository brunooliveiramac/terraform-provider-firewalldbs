package firewalldbs

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/core/service"
	"terraform-provider-firewalldbs/firewalldbs/data_provider/model"
)

func resourceCloseFirewall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloseFirewallCreate,
		ReadContext:   resourceCloseFirewallRead,
		UpdateContext: resourceCloseFirewallUpdate,
		DeleteContext: resourceCloseFirewallDelete,
		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AGENT_IP", ""),
			},
		},
	}
}

func resourceCloseFirewallCreate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	connection := providerConfig.(*model.Connection)

	serverName := resource.Get("server_name").(string)
	resourceGroup := resource.Get("resource_group_name").(string)
	serverId := resource.Get("server_id").(string)

	firewallRule := entity.ServerFirewallIpRule{
		IP:            connection.AgentIP,
		ServerName:    serverName,
		ResourceGroup: resourceGroup,
		Subscription:  connection.Subscription,
		ServerID:      serverId,
	}

	provider := service.GetProvider()

	err := provider.RemoveIp(&firewallRule, connection.Token)

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

	return diagnostics
}

func resourceCloseFirewallUpdate(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	connection := providerConfig.(*model.Connection)

	serverName := resource.Get("server_name").(string)
	resourceGroup := resource.Get("resource_group_name").(string)
	serverId := resource.Get("server_id").(string)

	firewallRule := entity.ServerFirewallIpRule{
		IP:            connection.AgentIP,
		ServerName:    serverName,
		ResourceGroup: resourceGroup,
		Subscription:  connection.Subscription,
		ServerID:      serverId,
	}

	provider := service.GetProvider()

	err := provider.RemoveIp(&firewallRule, connection.Token)

	if err != nil {
		msg := fmt.Sprintf("%s", err)

		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to Remove Agent IP",
			Detail:   msg,
		})

		return diagnostics
	}

	resource.SetId("AgentIP")

	return resourceCloseFirewallRead(ctx, resource, providerConfig)
}

func resourceCloseFirewallDelete(ctx context.Context, resource *schema.ResourceData, providerConfig interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics

	resource.SetId("")

	diagnostics = append(diagnostics, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Info",
		Detail:   "The deletion does not actually remove the firewall IP. To do it use the close resource.",
	})

	return diagnostics
}
