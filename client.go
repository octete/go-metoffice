/* This is a copy for inspiration from
 * github.com/sethvargo/go-fastly
 */
package metoffice

import (
	"fmt"
	"runtime"
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
}
