/* This is a copy for inspiration from
 * github.com/sethvargo/go-fastly
 */
package metoffice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/ajg/form"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

// APIKeyEnvVar is the name of the environment variable where the Metoffice API
// key should be read from.
const APIKeyEnvVar = "METOFFICE_API_KEY"

// DefaultEndpoint is the default endpoint for the metoffice.
const DefaultEndpoint = "http://datapoint.metoffice.gov.uk"

// ProjectURL is the url for this library.
var ProjectURL = "github.com/octete/go-metoffice"

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
func (c *Client) Put(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("PUT", p, ro)
}

// PutForm issues an HTTP PUT request with the given interface form-encoded
func (c *Client) PutForm(p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	return c.RequestForm("PUT", p, i, ro)
}

// Delete issues an HTTP DELETE request
func (c *Client) Delete(p string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("DELETE", p, ro)
}

// Request makes an HTTP requet against the HTTPClient using the given verb,
// Path, and request options.
func (c *Client) Request(verb, p string, ro *RequestOptions) (*http.Response, error) {
	req, err := c.RawRequest(verb, p, ro)
	if err != nil {
		return nil, err
	}

	resp, err := checkResp(c.HTTPClient.Do(req))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

//RequestForm makes an HTTP request with the give interface being endcoded as
// as form data.
func (c *Client) RequestForm(verb, p string, i interface{}, ro *RequestOptions) (*http.Response, error) {
	if ro == nil {
		ro = new(RequestOptions)
	}

	if ro.Headers == nil {
		ro.Headers = make(map[string]string)
	}
	ro.Headers["Content-Type"] = "application/x-www-form-urlencoded"

	buf := new(bytes.Buffer)
	if err := form.NewEncoder(buf).DelimitWith('|').Encode(i); err != nil {
		return nil, err
	}
	body := buf.String()

	ro.Body = strings.NewReader(body)
	ro.BodyLength = int64(len(body))

	return c.Request(verb, p, ro)
}

// checkResp wraps an HTTP request from the default client and verifies the
// request was successful. A non-200 request returns an error formatted to
// included any validation problems or otherwise.
func checkResp(resp *http.Response, err error) (*http.Response, error) {
	// If the err is already there, there was an error higher up the chain, so
	// just return that.
	if err != nil {
		return resp, err
	}

	switch resp.StatusCode {
	case 200, 201, 202, 204, 205, 206:
		return resp, nil
	default:
		return resp, NewHTTPError(resp)
	}
}

// decodeJSON is used to decode an HTTP response body into an interface as JSON
// TODO Check if this is the right way to do this.
func decodeJSON(out interface{}, body io.ReadCloser) error {
	defer body.Close()

	dec := json.NewDecoder(body)
	if err := dec.Decode(&out); err != nil {
		return err
	}

	// All good until here?
	return nil
}
