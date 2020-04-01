package awxcommunicator

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

//AWXCommunicator -
type AWXCommunicator struct {
	EndpointBase string
	Username     string
	Password     string
}

// basicAuth
func (c *AWXCommunicator) basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//DoRequest - Create request to AWX
func (c *AWXCommunicator) DoRequest(endpointURL string) string {
	request, _ := http.NewRequest("GET", (c.EndpointBase + endpointURL), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+c.basicAuth(c.Username, c.Password))

	// Pro test vypnuta kontrola certifikatu !!!
	client := &http.Client{

		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	}

	data, _ := ioutil.ReadAll(response.Body)
	return string(data)

}
