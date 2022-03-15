package data_provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider/model"
)

type Azure struct {}

func NewAzureProvider() core.Provider {
	return &Azure{}
}

func (a Azure) Login(credential *entity.Credential) (token string, err error) {

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

	resp, _ := client.Do(req)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyString := string(bodyBytes)

	msg := fmt.Sprintf("Request: %s, Response %s", string(requestDump), bodyString)

	if resp.StatusCode != 200 {
		log.Printf("Error : %d", resp.StatusCode)
		return "", errors.New(msg)
	}

	defer resp.Body.Close()

	responseModel := &model.LoginResponse{}

	err = json.Unmarshal(bodyBytes, &responseModel)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return "", errors.New(err.Error())
	}

	return responseModel.AccessToken, nil
}

func (a Azure) GetAgentIp() (ip string, err error) {

	requestUrl := fmt.Sprintf("https://ipinfo.io/ip")

	req, _ := http.NewRequest("GET", requestUrl, nil)

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

	if resp.StatusCode != 200 {
		log.Printf("Error : %d", resp.StatusCode)
		return "", errors.New(msg)
	}

	return bodyString, nil
}
