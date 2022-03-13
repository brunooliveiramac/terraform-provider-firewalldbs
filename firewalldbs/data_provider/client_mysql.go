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
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{Timeout: time.Second * 180}

type Connection struct {
	Subscription string
	Token        string
	AgentIP      string
}

type Credential struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Resource     string `json:"resource"`
	Tenant       string `json:"tenant_id"`
}

type Properties struct {
	Start string `json:"startIpAddress"`
	End   string `json:"endIpAddress"`
}

type AgentRequest struct {
	Properties Properties `json:"properties"`
}

type LoginResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
}

type ServerFirewallIpRule struct {
	IP            string `json:"ip"`
	ServerName    string `json:"server"`
	ResourceGroup string `json:"resource_group"`
	Subscription  string `json:"subscription"`
}

type FirewallRuleResponse struct {
	Properties Properties `json:"properties"`
	Name        string    `json:"name"`
}

func Login(credential *Credential) (token string, err error) {

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", credential.ClientId)
	data.Set("client_secret", credential.ClientSecret)
	data.Set("resource", "https://management.azure.com")

	requestUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", credential.Tenant)

	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	requestDump, err := httputil.DumpRequest(req, true)

	if err != nil {
		return "", err
	}

	fmt.Println(string(requestDump))

	resp, _ := client.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	println(bodyString)

	msg := fmt.Sprintf("Request: %s, Response %s", string(requestDump), bodyString)

	if resp.StatusCode != 200 {
		log.Printf("Error : %d", resp.StatusCode)
		return "", errors.New(msg)
	}

	defer resp.Body.Close()

	responseModel := &LoginResponse{}

	err = json.Unmarshal(bodyBytes, &responseModel)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return "", errors.New(err.Error())
	}

	fmt.Println(responseModel.AccessToken)

	return responseModel.AccessToken, nil
}

func GetFirewallRule(firewall *ServerFirewallIpRule, token string) (ruleName string, err error) {

	name := strings.Replace(firewall.IP, ".", "_", -1)

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBForMySQL/servers/%s/firewallRules/AllowAgent%s?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName, name)

	req, _ := http.NewRequest("GET", requestUrl, nil)

	tokenFmt := fmt.Sprintf("Bearer %s", token)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", tokenFmt)

	resp, _ := client.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	requestDump, err := httputil.DumpRequest(req, true)

	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	msg := fmt.Sprintf("Request: %s, Response %s", string(requestDump), bodyString)

	responseModel := &FirewallRuleResponse{}

	err = json.Unmarshal(bodyBytes, &responseModel)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return "", err
	}

	if resp.StatusCode != 200 {
		log.Printf("Error : %d", resp.StatusCode)
		return "", errors.New(msg)
	}

	println(responseModel.Name)

	return responseModel.Name, nil

}

func AddAgentIp(firewall *ServerFirewallIpRule, token string) (err error) {

	properties := Properties{
		Start: firewall.IP,
		End:   firewall.IP,
	}

	request := AgentRequest{properties}

	jsonValue, _ := json.Marshal(request)

	name := strings.Replace(firewall.IP, ".", "_", -1)

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBForMySQL/servers/%s/firewallRules/AllowAgent%s?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName, name)

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

	println(bodyString)

	msg := fmt.Sprintf("Request: %s, Response %s", string(requestDump), bodyString)

	if resp.StatusCode != 202 {
		log.Printf("Error : %d", resp.StatusCode)
		return errors.New(msg)
	}

	defer resp.Body.Close()

	return nil
}

func DeleteAgentIp(firewall *ServerFirewallIpRule, token string) (err error) {

	name := strings.Replace(firewall.IP, ".", "_", -1)

	requestUrl := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBForMySQL/servers/%s/firewallRules/AllowAgent%s?api-version=2017-12-01", firewall.Subscription, firewall.ResourceGroup, firewall.ServerName, name)

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

	println(bodyString)

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
