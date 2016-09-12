//
// Copyright (c) 2015-2016 Red Hat, Inc. / Nathan Sullivan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Credit: https://github.com/oVirt/ovirt-engine-sdk-ruby/blob/master/sdk/lib/ovirtsdk4/http.rb

package http

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/CpuID/ovirt-engine-sdk-go/sdk/slice"
	"github.com/CpuID/ovirt-engine-sdk-go/sdk/version"
)

// This type represents an HTTP request.
type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Query   map[string]string
	Body    string
}

// This type represents an HTTP response.
type Response struct {
	Body    string
	Code    int
	Headers map[string]string
	// Not used?
	// Message string
}

// This type (and its attached functions) are responsible for managing an HTTP connection to the engine server.
// It is intended as the entry point for the SDK, and it provides access to the `system` service and, from there,
// to the rest of the services provided by the API.
type Connection struct {
	url      url.URL
	username string
	password string
	token    string
	insecure bool
	caFile   string
	kerberos bool
	timeout  uint8
	compress bool
	//
	client *http.Client
}

// Creates a new connection to the API server.
func (c *Connection) Setup(input_url string, username string, password string, token string, insecure bool, ca_file string, kerberos bool, timeout uint8, compress bool) error {
	// Get the values of the parameters and assign default values:
	c.username = username
	c.password = password
	c.token = token
	c.insecure = insecure
	c.caFile = ca_file
	c.kerberos = kerberos
	c.timeout = timeout
	c.compress = compress

	// Check mandatory parameters:
	if len(input_url) == 0 {
		return errors.New("The 'input_url' parameter is mandatory.")
	}
	// TODOLATER: remove once kerberos is implemented
	if c.kerberos == true {
		return errors.New("Kerberos is not currently implemented.")
	}

	// Save the URL:
	use_url, err := url.Parse(input_url)
	if err != nil {
		return err
	}
	c.url = *use_url

	// Create the HTTP client:
	var disable_compress bool
	if compress == true {
		disable_compress = false
	} else {
		disable_compress = true
	}
	c.client = &http.Client{
		Timeout: time.Duration(timeout),
		Transport: &http.Transport{
			DisableCompression: disable_compress,
		},
	}
	if c.url.Scheme == "https" {
		c.client.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: insecure,
		}
		if len(c.ca_file) > 0 {
			// Check if the CA File specified exists.
			if _, err := os.Stat(c.caFile); os.IsNotExist(err) {
				return fmt.Errorf("The CA File '%s' doesn't exist.", c.caFile)
			}
			pool := x509.NewCertPool()
			ca_certs, err := ioutil.ReadAll(c.caFile)
			if err != nil {
				return err
			}
			if ok := pool.AppendCertsFromPEM([]byte{ca_certs}); ok == false {
				return fmt.Errorf("Failed to parse CA Certificate in file '%s'", c.caFile)
			}
			c.client.TLSClientConfig.RootCAs = pool
		}
	}

	// Debug output handled via GODEBUG env var. See https://golang.org/pkg/net/http/ for details.
}

// Returns the base URL of this connection.
func (c *Connection) Url() string {
	return c.url.String()
}

// Returns a reference to the root of the services tree.
func (c *Connection) SystemService() *SystemService {
	// TODO: implement
}

// Returns a reference to the service corresponding to the given path. For example, if the `path` parameter
// is `vms/123/diskattachments` then it will return a reference to the service that manages the disk
// attachments for the virtual machine with identifier `123`.
func (c *Connection) Service(path string) *Service {
	// TODO: implement
}

// Sends an HTTP request and waits for the response.
func (c *Connection) Send(r *Request) (*Response, error) {
	var result Response

	// Check if we already have an SSO access token:
	c.token = c.getAccessToken()

	// Build the URL:
	use_url := c.buildUrl(r.Path, r.Query)

	// Validate the method selected:
	if slice.stringInSlice(r.Method, []string{"DELETE", "GET", "PUT", "HEAD", "POST"}) == false {
		return &result, fmt.Errorf("The HTTP method '%s' is invalid, we expected one of DELETE/GET/PUT/HEAD/POST.", r.Method)
	}

	// Build the net/http request:
	req, err := http.NewRequest(r.Method, use_url, nil)
	if err != nil {
		return &result, err
	}

	// Add request headers:
	for k1, v1 := range r.Headers {
		req.Header.Add(k1, v1)
	}
	req.Header.Add("User-Agent", fmt.Sprintf("GoSDK/%s", version.SdkVersion))
	req.Header.Add("Version", "4")
	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	// Send the request and wait for the response:
	resp, err := c.client.Do(req)
	if err != nil {
		return &result, err
	}

	// Return the response:
	defer resp.Body.Close()
	result.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return &result, err
	}
	result.code = resp.StatusCode
	result.headers = make(map[string]string)
	for k2, v2 := range resp.Headers {
		result.headers[k2] = v2
	}
	return result, nil
}

// Obtains the access token from SSO to be used for bearer authentication.
func (c *Connection) getAccessToken() (string, error) {
	// Build the URL and parameters required for the request:
	url, parameters := c.buildSsoAuthRequest()

	// Send the response and wait for the request:
	response := c.getSsoResponse(url, parameters)

	// Top level array already handled in getSsoResponse() generically.

	if len(response.ssoError) > 0 {
		return "", fmt.Errorf("Error during SSO authentication %s: %s", response.ssoErrorCode, response.ssoError)
	}

	return response.accessToken, nil
}

// Revoke the SSO access token.
func (c *Connection) revokeAccessToken() error {
	// Build the URL and parameters required for the request:
	url, parameters := c.buildSsoRevokeRequest()

	// Send the response and wait for the request:
	response := c.getSsoResponse(url, parameters)

	// Top level array already handled in getSsoResponse() generically.

	if len(response.ssoError) > 0 {
		return fmt.Errorf("Error during SSO revoke %s: %s", response.ssoErrorCode, response.ssoError)
	}
	return nil
}

type ssoResponseJsonParent struct {
	children []ssoResponseJson
}

type ssoResponseJson struct {
	accessToken  string `json:"access_token"`
	ssoError     string `json:"error"`
	ssoErrorCode string `json:"error_code"`
}

// Execute a get request to the SSO server and return the response.
func (c *Connection) getSsoResponse(input_url string, parameters map[string]string) (ssoResponseJson, error) {
	// Create the HTTP client handle for SSO:
	client = &http.Client{
		Timeout: time.Duration(c.timeout),
	}

	// Configure TLS parameters:
	if input_url.Scheme == "https" {
		client.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: c.insecure,
		}
		if len(c.ca_file) > 0 {
			// Check if the CA File specified exists.
			if _, err := os.Stat(c.caFile); os.IsNotExist(err) {
				return ssoResponseJson{}, fmt.Errorf("The CA File '%s' doesn't exist.", c.caFile)
			}
			pool := x509.NewCertPool()
			ca_certs, err := ioutil.ReadAll(c.caFile)
			if err != nil {
				return ssoResponseJson{}, err
			}
			if ok := pool.AppendCertsFromPEM([]byte{ca_certs}); ok == false {
				return ssoResponseJson{}, fmt.Errorf("Failed to parse CA Certificate in file '%s'", c.caFile)
			}
			client.TLSClientConfig.RootCAs = pool
		}
	}

	// Configure authentication:
	// TODOLATER: implement, skipped for now. kerberos support in Golang seems "interesting"?
	// if c.kerberos == true {
	// }

	// POST request body:
	var body_values url.Values
	for k1, v1 := range parameters {
		body_values[k1] = []string{v1}
	}

	// Build the net/http request:
	req, err := http.NewRequest("POST", input_url, body_values.Encode())
	if err != nil {
		return err
	}

	// Add request headers:
	req.Header.Add("User-Agent", fmt.Sprintf("GoSDK/%s", version.SdkVersion))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	// Send the request and wait for the response:
	resp, err := client.Do(req)
	if err != nil {
		return ssoResponseJson{}, err
	}
	defer resp.Body.Close()

	// Parse and return the JSON response:
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ssoResponseJson{}, err
	}
	var json1 ssoResponseJson
	json1, err1 := json.Unmarshal(body, &result1)
	if err1 != nil {
		// Maybe it's array encapsulated, try the other approach.
		var json2 ssoResponseJsonParent
		json2, err2 := json.Unmarshal(body, &result2)
		if err2 != nil {
			return fmt.Errorf("Errors for both JSON unmarshal methods (array/non-array) for SSO response: %s / %s", err1.Error(), err2.Error())
		}
		if len(json2) == 0 {
			return errors.New("SSO response is a zero length array.")
		}
		json1.accessToken = json2.children[0].accessToken
		json1.ssoError = json2.children[0].ssoError
		return json1, nil
	}
	return json1, nil
}

// Builds a the URL and parameters to acquire the access token from SSO.
func (c *Connection) buildSsoAuthRequest() (string, map[string]string) {
	// Compute the entry point and the parameters:
	parameters := map[string]string{
		"scope": "ovirt-app-api",
	}

	var entry_point string
	if c.kerberos == true {
		entry_point = "token-http-auth"
		parameters["grant_type"] = "urn:ovirt:params:oauth:grant-type:http"
	} else {
		entry_point = "token"
		parameters["grant_type"] = "password"
		parameters["username"] = c.username
		parameters["password"] = c.password
	}

	// Compute the URL:
	sso_url := c.url
	sso_url.Path = fmt.Sprintf("/ovirt-engine/sso/oauth/%s", entry_point)

	// Return the URL and the parameters:
	return sso_url.String(), parameters
}

// Builds a the URL and parameters to revoke the SSO access token.
// string = the URL of the SSO service
// map = hash containing the parameters required to perform the revoke
func (c *Connection) buildSsoRevokeRequest() (string, map[string]string) {
	// Compute the parameters:
	parameters := map[string]string{
		"scope": "",
		"token": c.token,
	}

	// Compute the URL:
	sso_url := c.url
	sso_url.Path = "/ovirt-engine/services/sso-logout"

	// Return the URL and the parameters:
	return sso_url.String(), parameters
}

// Tests the connectivity with the server. If connectivity works correctly it returns a nil error. If there is any
// connectivity problem it will return an error containing the reason as the message.
func (c *Connection) Test() error {
	return c.SystemService.Get()
}

// Performs the authentication process and returns the authentication token. Usually there is no need to
// call this method, as authentication is performed automatically when needed. But in some situations it
// may be useful to perform authentication explicitly, and then use the obtained token to create other
// connections, using the `token` parameter of the constructor instead of the user name and password.
func (c *Connection) Authenticate() {
	c.token = c.getAccessToken()
}

// Indicates if the given object is a link. An object is a link if it has an `href` attribute.
func (c *Connection) IsLink(object string) bool {
	// TODO: implement
}

// Follows the `href` attribute of the given object, retrieves the target object and returns it.
func (c *Connection) FollowLink(object string) error {
	// TODO: implement
}

// Releases the resources used by this connection.
func (c *Connection) Close() {
	if len(token) > 0 {
		c.revokeAccessToken()
	}
}

// Builds a request URL from a path, and the set of query parameters.
func (c *Connection) BuildUrl(path string, query map[string]string) string {
	url = fmt.Sprintf("%s%s", c.url.String(), path)
	if len(query) > 0 {
		var values url.Values
		for k, v := range query {
			values[k] = []string{v}
		}
		url = fmt.Sprintf("%s?%s", url, values.Encode())
	}
	return url
}
