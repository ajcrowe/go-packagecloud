package packagecloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Package Cloud API URL
const BaseURL = "https://packagecloud.io/api/v1"

type Client struct {
	Token  string
	client *http.Client
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		token = os.Getenv("PACKAGECLOUD_TOKEN")
		if token == "" {
			return nil, errors.New("PACKAGECLOUD_TOKEN unset")
		}
	}
	httpClient := http.DefaultClient
	return &Client{
		Token:  token,
		client: httpClient,
	}, nil
}

func (c *Client) do(req *http.Request, status int, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	// check status code is what is expected
	if resp.StatusCode != status {
		return resp, errors.New("Incorrect status code returned")
	}

	// decode resp body into struct
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

func createUri(section, user, repo, resource string) *url.URL {
	strUrl := fmt.Sprintf("%s/%s/%s/%s/%s", BaseURL, section, user, repo, resource)
	reqUrl, _ := url.Parse(strUrl)
	return reqUrl
}
