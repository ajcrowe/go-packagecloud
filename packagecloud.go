package packagecloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
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

// wrapper function for http.NewRequest to add access token
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	req.SetBasicAuth(c.Token, "")
	return req, err
}

func (c *Client) do(req *http.Request, status int, v interface{}) (*http.Response, error) {
	dump, err := httputil.DumpRequest(req, true)
	fmt.Println(string(dump))

	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	dump, err = httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))

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
