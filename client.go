// Package gomarathon provIDes a client to interact with a marathon
// api. on http or https
package gomarathon

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is containing the configured http.Client
// and the host url
type HttpBasicAuth struct {
	User string
	Pass string
}

type Client struct {
	Url        string
	HTTPClient *http.Client
	Auth       *HttpBasicAuth
}

type UpdateResp struct {
	DeploymentID string `json:"deploymentId"`
}

// Actual version of the marathon api
const (
	APIVersion = "/v2"
)

// NewClient return a pointer to the new client
func NewClient(host string, auth *HttpBasicAuth, tlsConfig *tls.Config) (*Client, error) {
	// ValIDate url
	h, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("can't parse host %s", host)
	}

	return &Client{
		Url:        h.String(),
		HTTPClient: newHTTPClient(h, tlsConfig),
		Auth:       auth,
	}, nil
}

// do the actual prepared request in request()
func (c *Client) do(method, path string, data interface{}) ([]byte, int, error) {
	var params io.Reader
	var resp *http.Response

	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, -1, err
		}
		params = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, c.Url+path, params)
	if err != nil {
		return nil, -1, err
	}

	// Prepare and do the request
	req.Header.Set("User-Agent", "gomarathon")
	req.Header.Set("Content-Type", "application/json")

	if c.Auth != nil {
		req.SetBasicAuth(c.Auth.User, c.Auth.Pass)
	}

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1, err
	}
	if resp.StatusCode >= 400 {
		return nil, resp.StatusCode, fmt.Errorf("%d: %s", resp.StatusCode, body)
	}

	return body, resp.StatusCode, nil
}

// request prepare the request by setting the correct methods and parameters
// TODO:
// 	- find a better way to build parameters
func (c *Client) request(options *RequestOptions) (*Response, error) {

	if options.Path == "" {
		options.Path = "apps"
	}

	if options.Method == "" {
		options.Method = "GET"
	}

	path := fmt.Sprintf("%s/%s", APIVersion, options.Path)

	if options.Params != nil {
		v := url.Values{}

		if options.Params.Cmd != "" {
			v.Set("cmd", options.Params.Cmd)
		}

		if options.Params.Host != "" {
			v.Set("host", options.Params.Host)
		}

		if options.Params.Scale {
			v.Set("scale", "true")
		}

		if options.Params.CallbackURL != "" {
			v.Set("CallbackURL", options.Params.CallbackURL)
		}

		path = fmt.Sprintf("%s?%s", path, v.Encode())
	}

	data, code, err := c.do(options.Method, path, options.Datas)
	if err != nil {
		return nil, err
	}
	resp := &Response{
		Code: code,
	}

	//updated
	if resp.Code == 200 {
		updateResp := UpdateResp{}
		err := json.Unmarshal(data, &updateResp)
		if err == nil {
			resp.DeploymentId = updateResp.DeploymentID
		} else {
			fmt.Println("Error unmashaling data response")
		}
		//created
	} else if resp.Code == 201 {
		app := Application{}
		err := json.Unmarshal(data, &app)
		if err == nil {
			resp.DeploymentId = app.Deployments[0].ID
		} else {
			fmt.Println("Error unmashaling data response")
		}
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
