package packagecloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	//"net/http/httputil"
	"net/url"
	"os"
)

// Package Cloud API URL
const (
	defaultBaseURL  = "https://packagecloud.io"
	defaultMimeType = "application/json"
	apiVersion      = "v1"
)

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

	// create our client struct
	httpClient := http.DefaultClient
	client := &Client{
		Token:  token,
		client: httpClient,
	}

	// load package distribution information
	_, err := client.loadPackageDistributions()
	if err != nil {
		return client, err
	}
	return client, nil
}

// wrapper function for http.NewRequest to add access token and custom mimetype
func (c *Client) NewRequest(method, urlStr, contentType string, body io.Reader) (*http.Request, error) {
	// set default contact type if not set
	if contentType == "" {
		contentType = defaultMimeType
	}
	// get new request
	req, err := http.NewRequest(method, urlStr, body)
	// set basic auth username to api token
	req.SetBasicAuth(c.Token, "")
	// set content type
	req.Header.Set("Content-Type", contentType)
	return req, err
}

func (c *Client) do(req *http.Request, status int, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	// check status code is what is expected
	if resp.StatusCode != status {
		return resp, errors.New("packagecloud: bad status code returned")
	}

	// decode resp body into struct
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

func createUri(section, user, repo, resource string) *url.URL {
	strUrl := fmt.Sprintf("%s/api/%s/%s/%s/%s/%s", defaultBaseURL, apiVersion, section, user, repo, resource)
	reqUrl, _ := url.Parse(strUrl)
	return reqUrl
}

func createUriFromPath(path string) *url.URL {
	strUrl := fmt.Sprintf("%s%s", defaultBaseURL, path)
	reqUrl, _ := url.Parse(strUrl)
	return reqUrl
}
