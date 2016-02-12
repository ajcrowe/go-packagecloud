package packagecloud

import (
	"fmt"
	"net/http"
)

// License struct see https://packagecloud.io/docs/api#resource_licenses
type License struct {
	License   string `json:"license"`
	Signature string `json:"signature"`
}

// GetLicense returns a License struct for the specified license key
func (c *Client) GetLicense(key string) (License, *http.Response, error) {
	var license License
	reqURL := createURIFromPath(fmt.Sprintf("/api/%s/licenses/%s/license.json", apiVersion, key))

	// Create HTTP request
	req, err := c.NewRequest("GET", reqURL.String(), "", nil)
	if err != nil {
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &license)
	if err != nil {
		return license, resp, err
	}
	return license, resp, nil
}
