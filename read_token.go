package packagecloud

import (
	//"time"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

type ReadToken struct {
	// Unique Id for the read token
	Id int `json:"id"`

	// Name of the read token
	Name string `json:"name"`

	// Read token value used for access
	Value string `json:"value"`
}

// ListReadTokens returns a slice of pointers to ReadToken structs which are
// accociated with the masterToken.
func (c *Client) ListReadTokens(user, repo, tokenUrl string) ([]*ReadToken, error) {
	var tokens []*ReadToken
	// Construct URL for request
	matched, _ := regexp.MatchString(`^/api/v1/repos/\w+/\w+/master_tokens/\d+$`, tokenUrl)
	if !matched {
		return tokens, errors.New("Invalid master token URL")
	}

	reqUrl := fmt.Sprint(tokenUrl + "/read_tokens.json")
	// Create HTTP request
	req, err := c.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return tokens, err
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &tokens)
	if err != nil {
		fmt.Printf("packagecloud: Error bad response code: %s", resp.StatusCode)
		return tokens, err
	}
	return tokens, nil
}

// CreateReadToken creates a new read token for the specified master token value.
func (c *Client) CreateReadToken(user, repo, tokenUrl, name string) (ReadToken, error) {
	var token ReadToken
	matched, _ := regexp.MatchString(`^/api/v1/repos/\w+/\w+/master_tokens/\d+$`, tokenUrl)
	if !matched {
		return token, errors.New("Invalid master token URL")
	}
	body := []byte(fmt.Sprintf("read_token[name]=%s", name))

	reqUrl := fmt.Sprint(tokenUrl + "/read_tokens.json")
	// Create HTTP request
	req, err := c.NewRequest("POST", reqUrl, bytes.NewReader(body))
	if err != nil {
		return token, err
	}

	// Do request
	resp, err := c.do(req, http.StatusCreated, &token)
	if err != nil {
		fmt.Printf("packagecloud: Error bad response code: %s", resp.StatusCode)
		return token, err
	}
	return token, nil
}

// DestroyReadToken deletes the read token with specified id from the master token
func (c *Client) DestroyReadToken(user, repo, masterToken string, id int) error {
	return nil
}
