/* This is a copy for inspiration from
 * github.com/sethvargo/go-fastly
 */
package metoffice

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// APIKeyEnvVar is the name of the environment variable where the Metoffice API
// key should be read from.
const APIKeyEnvVar = "METOFFICE_API_KEY"

// DefaultEndpoint is the default endpoint for the metoffice.
const DefaultEndpoint = "http://datapoint.metoffice.gov.uk"

// ProjectURL is the url for this library.
var ProjectURL = "gitlab.com/octete/metoffice-go"

// ProjectVersion is the version of this library
var ProjectVersion = "0.1"

// UserAgent is the user agent for this particular client.
var UserAgent = fmt.Sprintf("MetofficeGo/%s (+%s; %s)",
	ProjectVersion, ProjectURL, runtime.Version())

// Client is the main entrypoint to the Metoffice golang API library.
type Client struct {
	// Address is the address of the Metoffice's API endpoint
	Address string

	// HTTPClient is the HTTP client to use. If one is not provided, a default
	// client will be used.
	HTTPClient *http.Client

	// apiKey is the Fastly API to authenticate requests.
	apiKey string

	// url is the parsed URL from Address
	url *url.URL
}

// DefaultClient instantiates a new Metoffice API client. This function requires
// the environment variable `METOFFICE_API_KEY` is set and contains a valid key
// to authenticate with the Metoffice.
func DefaultClient() *Client {
	client, err := NewClient(os.Getenv(APIKeyEnvVar))
	if err != nil {
		panic(err)
	}
	return client
}

// NewClient creates a new API client with the given key.
func NewClient(key string) (*Client, error) {
	client := &Client{apiKey: key}
	return client.init()
}

func (c *Client) init() (*Client, error) {
	if len(c.Address) == 0 {
		c.Address = DefaultEndpoint
	}

	u, err := url.Parse(c.Address)
	if err != nil {
		return nil, err
	}
	c.url = u

	if c.HTTPClient == nil {
		c.HTTPClient = cleanhttp.DefaultClient()
	}

	return c, nil
}

// Get issues an HTTP GET request
func (c *Client) Get(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("GET", p, ro)
}

// Head issues an HTTP HEAD request
func (c *Client) Head(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("HEAD", p, ro)
}

// Post issues an HTTP POST request
func (c *Client) Post(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("POST", p, ro)
}

// PostForm issues an HTTP POST request with the given interface form-encoded
func (c *Client) PostForm(p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	return c.RequestForm("POST", p, i, ro)
}

// Post issues an HTTP PUT request
func (c *Client) Get(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("PUT", p, ro)
}

// PutForm issues an HTTP PUT request with the given interface form-encoded
func (c *Client) PutForm(p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	return c.RequestForm("PUT", p, i, ro)
}

// Delete issues an HTTP DELETE request
func (c *Client) Get(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("DELETE", p, ro)
}

// Request makes an HTTP requet against the HTTPClient using the given verb,
// Path, and request options.
func (c *Client) Request(verb, p string, ro *RequestOptions) (*http.Response, error) {
	req, err := c.RawRequest(verb, p, ro)
	if err != nil {
		return nil, err
	}

	resq, err := checkResp(c.HTTPClient.Do(req))
	if err != nil {
		return resp, err
	}

	return resp, nil
}
