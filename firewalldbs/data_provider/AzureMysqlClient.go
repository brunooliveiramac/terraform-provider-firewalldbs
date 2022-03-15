package data_provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider/model"
	"time"
)

var client = &http.Client{Timeout: time.Second * 180}

type MysqlProvider struct {}

func NewMysqlProvider() core.Database {
	return &MysqlProvider{}
}

func (m MysqlProvider) AddAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {

	properties := model.Properties{
		Start: firewall.IP,
		End:   firewall.IP,
	}

	request := model.AgentRequest{Properties: properties}

	jsonValue, _ := json.Marshal(request)

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBForMySQL/servers/%s/firewallRules/AllowAgent?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName)

	req, _ := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(jsonValue))

	tokenFmt := fmt.Sprintf("Bearer %s", token)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenFmt)

	resp, _ := client.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	requestDump, err := httputil.DumpRequest(req, true)

	if err != nil {
		return err
	}

	bodyString := string(bodyBytes)

	msg := fmt.Sprintf("Request: %s, Response %s", string(requestDump), bodyString)

	if resp.StatusCode != 202 {
		log.Printf("Error : %d", resp.StatusCode)
		return errors.New(msg)
	}

	defer resp.Body.Close()

	return nil
}


func (m MysqlProvider) DeleteAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBForMySQL/servers/%s/firewallRules/AllowAgent?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName)

	req, _ := http.NewRequest("DELETE", requestUrl, nil)

	tokenFmt := fmt.Sprintf("Bearer %s", token)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenFmt)

	resp, _ := client.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	bodyString := string(bodyBytes)

	requestDump, err := httputil.DumpRequest(req, true)

	if err != nil {
		return err
	}

	msg := fmt.Sprintf("Request: %s, Response %s", string(requestDump), bodyString)

	if resp.StatusCode != 202 {
		log.Printf("Error : %d", resp.StatusCode)
		return errors.New(msg)
	}

	return nil
}
