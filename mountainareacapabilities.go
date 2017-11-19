package metoffice

import (
	"fmt"
	"time"
)

// mountainforecast list

// MountainForecastItem has the data for each of the forecast areas that the metoffice
// provides, with validity data and destination URL.
type MountainForecastItem struct {
	DataDate   time.Time `json:"DataDate"`
	ValidFrom  time.Time `json:"ValidFrom"`
	ValidTo    time.Time `json:"ValidTo"`
	CreateDate time.Time `json:"CreateDate"`
	URI        string    `json:"URI"`
	Area       string    `json:"Area"`
	Risk       string    `json:"Risk"`
}

// MountainForecastList is a struct that has a list of mountain
// areas (MountainForecastItem)
type MountainForecastList struct {
	MF []MountainForecastItem `json:"MountainForecast"`
}

// MountainForecastListInput is the input we get from the metoffice
type MountainForecastListInput struct {
	MFL MountainForecastList `json:"MountainForecastList"`
}

// GetMountainAreaCapabilities returns all the mountain areas available from the Metoffice
// in an array.
// The mountain area forecast capabilities data feed provides a summary of which results are available from the get
// mountain area forecast by site ID data feed, specifying the creation dates, valid from and to dates, and the general
// risk for each mountain area.
// func (c *Client) ListMountainForecasts() ([]MountainForecastItem, error) {
func (c *Client) GetMountainAreaCapabilities() ([]MountainForecastItem, error) {

	path := fmt.Sprintf("/public/data/txt/wxfcs/mountainarea/%s/capabilities", "json")
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var b *MountainForecastListInput
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}
	var mfl *MountainForecastList
	mfl = &b.MFL

	return mfl.MF, nil
}
