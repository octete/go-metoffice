package metoffice

import (
	"fmt"
)

// MountainAreaSitelist represents the list of mountain areas
// in the metoffice.
type MountainAreaSitelist struct {
	Locations Locations `json:"Locations"`
}

// Locations indicates the list of locations
type Locations struct {
	Locations []Location `json:"Location"`
}

// Location is each of the locations.
type Location struct {
	ID   string `json:"@Id"`
	Name string `json:"@Name"`
}

// GetMountainAreaSitelist Returns a list of locations the mountain area
// forecast data feed provides data for .You can use this to find the
// ID of the site that you are interested in.
func (c *Client) GetMountainAreaSitelist() (*Locations, error) {

	// TODO template json as a variable.
	path := fmt.Sprintf("/public/data/txt/wxfcs/mountainarea/%s/sitelist", "json")

	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var b *MountainAreaSitelist
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}

	return &b.Locations, nil
}
