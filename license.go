package packagecloud

import (
	"fmt"
	"net/http"
)

type License struct {
	License   string `json:"license"`
	Signature string `json:"signature"`
}

func (c *Client) GetLicense(name string) (License, *http.Response, error) {
	var license License
	reqUrl := createUriFromPath(fmt.Sprintf("/api/%s/licenses/%s/license.json", apiVersion, name))

	// Create HTTP request
	req, err := c.NewRequest("GET", reqUrl.String(), "", nil)
	if err != nil {
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &license)
	if err != nil {
		return license, resp, err
	}
	return license, resp, nil
}
