package packagecloud

import (
	"fmt"
	"net/http"
)

// Distribution struct see https://packagecloud.io/docs/api#resource_distributions
type Distribution struct {
	DisplayName string                `json:"display_name"`
	IndexName   string                `json:"index_name"`
	Versions    []DistributionVersion `json:"versions"`
}

// DistributionVersion struct contains the nested version data from the Distribution
type DistributionVersion struct {
	ID            int    `json:"id"`
	DisplayName   string `json:"display_name"`
	IndexName     string `json:"index_name"`
	VersionNumber string `json:"version_number"`
}

// Distributions is a slice of Distribution
type Distributions []Distribution

var (
	distroURI      = fmt.Sprintf("/api/%s/distributions.json", apiVersion)
	packageDistros = map[string]Distributions{}
	packageTypes   = []string{"deb", "rpm", "gem", "dsc"}
)

func (c *Client) loadPackageDistributions() (*http.Response, error) {
	reqURL := createURIFromPath(distroURI)

	// Create HTTP request
	req, err := c.NewRequest("GET", reqURL.String(), "", nil)
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

// GetDistributions returns a slice of all known Distributions
func (c *Client) GetDistributions() map[string]Distributions {
	return packageDistros
}
