package gomarathon

import (
	"fmt"
)

/*
* Application requests
 */

func (c *Client) ListApps() (*Response, error) {
	return c.ListAppsByCmd("")
}

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

func (c *Client) ListAppVersions(appId string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/versions", appId),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetApp(appId string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s", appId),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("Request Error")
	}
	return r, nil
}

func (c *Client) GetAppVersion(appId string, version string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/versions/%s", appId, version),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("Request Error")
	}
	return r, nil
}

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

func (c *Client) UpdateApp(appId string, app *Application) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appId),
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

func (c *Client) DeleteApp(appId string) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appId),
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
