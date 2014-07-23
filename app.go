package gomarathon

import (
	"fmt"
)

/*
* Application requests
 */

// ListApps is listing all apps
func (c *Client) ListApps() (*Response, error) {
	return c.ListAppsByCmd("")
}

// ListAppsByCmd list all apps by cmd filter
func (c *Client) ListAppsByCmd(filter string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps"),
		Params: &Parameters{
			Cmd: filter,
		},
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// ListAppVersions list app versions from a define appid
func (c *Client) ListAppVersions(appID string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/versions", appID),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetApp gets a single app with production version
func (c *Client) GetApp(appID string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s", appID),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("request rrror")
	}
	return r, nil
}

// GetAppVersion get a single version from a single app
func (c *Client) GetAppVersion(appID string, version string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/versions/%s", appID, version),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("request error")
	}
	return r, nil
}

// CreateApp Create a new Application
func (c *Client) CreateApp(app *Application) (*Response, error) {
	// TODO : VALIDATE DATAS
	options := &RequestOptions{
		Path:   "apps",
		Datas:  app,
		Method: "POST",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 201 {
			return r, nil
		}
	}
	return nil, err
}

// UpdateApp update the app but thoses changes are made for the next running app and does
// not shut down the production applications
func (c *Client) UpdateApp(appID string, app *Application) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Datas:  app,
		Method: "PUT",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err

}

// DeleteApp delete this app from the cluster
func (c *Client) DeleteApp(appID string) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Method: "DELETE",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}
