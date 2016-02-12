package packagecloud

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

const (
	masterTokenRegex = `^/api/v1/repos/\w+/\w+/master_tokens/\d+$`
)

var errMasterTokenRegex = fmt.Errorf("invalid master token url should match %s", masterTokenRegex)

// ReadToken struct see https://packagecloud.io/docs/api#resource_read_tokens
type ReadToken struct {
	// Unique ID for the read token
	ID int `json:"id"`

	// Name of the read token
	Name string `json:"name"`

	// Read token value used for access
	Value string `json:"value"`
}

// ListReadTokens returns a slice of pointers to ReadToken structs which are
// accociated with the masterToken.
func (c *Client) ListReadTokens(user, repo, tokenURL string) ([]*ReadToken, *http.Response, error) {
	var tokens []*ReadToken
	// Construct URL for request
	matched, _ := regexp.MatchString(masterTokenRegex, tokenURL)
	if !matched {
		return tokens, nil, errMasterTokenRegex
	}

	reqURL := fmt.Sprint(tokenURL + "/read_tokens.json")
	// Create HTTP request
	req, err := c.NewRequest("GET", reqURL, "", nil)
	if err != nil {
		return tokens, nil, err
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &tokens)
	if err != nil {
		return tokens, resp, err
	}
	return tokens, resp, nil
}

// CreateReadToken creates a new read token for the specified master token value.
func (c *Client) CreateReadToken(user, repo, tokenURL, name string) (ReadToken, *http.Response, error) {
	var token ReadToken
	matched, _ := regexp.MatchString(masterTokenRegex, tokenURL)
	if !matched {
		return token, nil, errMasterTokenRegex
	}
	body := []byte(fmt.Sprintf("read_token[name]=%s", name))

	reqURL := createUriFromPath(fmt.Sprint(tokenURL + "/read_tokens.json"))
	// Create HTTP request
	req, err := c.NewRequest("POST", reqURL.String(), "multipart/form-data", bytes.NewReader(body))
	if err != nil {
		return token, nil, err
	}

	// Do request
	resp, err := c.do(req, http.StatusCreated, &token)
	if err != nil {
		return token, resp, err
	}
	return token, resp, nil
}

// DestroyReadToken deletes the read token with specified id from the master token
func (c *Client) DestroyReadToken(user, repo, masterToken string, id int) (*http.Response, error) {
	return nil, nil
}
