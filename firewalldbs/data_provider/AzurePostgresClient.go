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
	"time"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider/model"
)

type PostgresProvider struct{}

func NewAzurePostgresProvider() core.Database {
	return &PostgresProvider{}
}

func (p PostgresProvider) AddAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {

	properties := model.Properties{
		Start: firewall.IP,
		End:   firewall.IP,
	}

	request := model.AgentRequest{Properties: properties}

	jsonValue, _ := json.Marshal(request)

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/servers/%s/firewallRules/AllowAgent?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName)

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

	headers := resp.Header
	asyncUrl := headers.Get("azure-asyncoperation")

	check, err2 := p.CheckAgentIpAllowed(asyncUrl, token)
	if err2 != nil {
		return err2
	}
	if !check {
		errMsg := fmt.Sprintf("Failed to add IP %s in PostgreSQL.", firewall.IP)
		return errors.New(errMsg)
	}

	defer resp.Body.Close()

	return nil
}

func (p PostgresProvider) DeleteAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/servers/%s/firewallRules/AllowAgent?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName)

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

func (p PostgresProvider) CheckAgentIpAllowed(url string, token string) (bool, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	tokenFmt := fmt.Sprintf("Bearer %s", token)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenFmt)

	client := &http.Client{}
	var response model.AsyncResponse

	startTime := time.Now()
	endTime := startTime.Add(5 * time.Minute)
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		resp, err := client.Do(req)
		if err != nil {
			ticker.Stop()
			return false, err
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			ticker.Stop()
			return false, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			ticker.Stop()
			return false, err
		}

		if response.Status == "Succeeded" {
			log.Printf("IP address added successfully.")
			ticker.Stop()
			return true, nil
		}

		log.Printf("IP address not yet added.")

		if time.Now().After(endTime) {
			ticker.Stop()
			break
		}
	}

	return false, nil
}