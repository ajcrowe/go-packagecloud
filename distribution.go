package packagecloud

import (
	"fmt"
	"net/http"
)

type Distribution struct {
	DisplayName string                `json:"display_name"`
	IndexName   string                `json:"index_name"`
	Versions    []DistributionVersion `json:"versions"`
}

type DistributionVersion struct {
	Id            int    `json:"id"`
	DisplayName   string `json:"display_name"`
	IndexName     string `json:"index_name"`
	VersionNumber string `json:"version_number"`
}

type Distributions []Distribution

var (
	distroUri      = fmt.Sprintf("/api/%s/distributions.json", apiVersion)
	packageDistros = map[string]Distributions{}
	packageTypes   = []string{"deb", "rpm", "gem", "dsc"}
)

func (c *Client) loadPackageDistributions() (*http.Response, error) {
	reqUrl := createUriFromPath(distroUri)

	// Create HTTP request
	req, err := c.NewRequest("GET", reqUrl.String(), "", nil)
	if err != nil {
		return nil, err
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &packageDistros)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (c *Client) GetDistributions() map[string]Distributions {
	return packageDistros
}
