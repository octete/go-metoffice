package metoffice

import (
	"fmt"
	"time"
)

// Go code to deal with what we get from the MetOffice
//type AutoGenerated struct {
//	MountainForecastList struct {
//		MountainForecast []struct {
//			DataDate time.Time `json:"DataDate"`
//			ValidFrom time.Time `json:"ValidFrom"`
//			ValidTo time.Time `json:"ValidTo"`
//			CreatedDate time.Time `json:"CreatedDate"`
//			URI string `json:"URI"`
//			Area string `json:"Area"`
//			Risk string `json:"Risk"`
//		} `json:"MountainForecast"`
//	} `json:"MountainForecastList"`
//}

// MountainForecast has the data for each of the forecast areas that the metoffice
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

// ListMountainForecasts returns all the mountain areas available from the Metoffice
// in an array.
func (c *Client) ListMountainForecasts() ([]MountainForecastItem, error) {
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

// MountainAreaInput is empty
type MountainAreaInput struct {
	Report Report `json:"report"`
}

type Report struct {
	CreatingAuthority string       `json:"creating-authority"`
	CreationTime      string       `json:"creation-time"`
	Title             string       `json:"title"`
	Location          string       `json:"location"`
	Issue             Issue        `json:"issue"`
	ValidFrom         time.Time    `json:"ValidFrom"`
	ValidTo           time.Time    `json:"ValidTo"`
	Validity          string       `json:"Validity"`
	IssuedDate        string       `json:"IssuedDate"`
	Hazards           Hazards      `json:"Hazard"`
	Overview          string       `json:"Overview"`
	ForecastDay0      ForecastDay0 `json:"Forecast_Day0"`
	ForecastDay1      ForecastDay1 `json:"Forecast_day1"`
	OutlookDay2       string       `json:"Outlook_Day2"`
	OutlookDay3       string       `json:"Outlook_Day3"`
	OutlookDay4       string       `json:"Outlook_Day4"`
}

type Issue struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type Hazards struct {
	Hazard []Hazard `json:"Hazard"`
}

type Hazard struct {
	No       string `json:"no"`
	Element  string `json:"Element"`
	Risk     string `json:"Risk"`
	Comments string `json:"Comments"`
}

type ForecastDay0 struct {
	Weather       string     `json:"Weather"`
	Visibility    string     `json:"Visibility"`
	HillFog       string     `json:"HillFog"`
	MaxWindLevel  string     `json:"MaxWindLevel"`
	MaxWind       string     `json:"MaxWind"`
	TempLowLevel  string     `json:"TempLowLevel"`
	TempHighLevel string     `json:"TempHighLevel"`
	FreezingLevel string     `json:"FreezingLevel"`
	WeatherPPN    WeatherPPN `json:"WeatherPPN"`
}

type WeatherPPN struct {
	WxPeriod []WxPeriod `json:"WxPeriod"`
}

type WxPeriod struct {
	Period      string `json:"period"`
	PeriodDesc  string `json:"Period"`
	Weather     int    `json:"Weather"`
	Probability string `json:"Probability"`
	PpnType     string `json:"Ppn_type"`
}

type ForecastDay1 struct {
	Weather       string `json:"Weather"`
	Visibility    string `json:"Visibility"`
	HillFog       string `json:"HillFog"`
	MaxWindLevel  string `json:"MaxWindLevel"`
	MaxWind       string `json:"MaxWind"`
	TempLowLevel  string `json:"TempLowLevel"`
	TempHighLevel string `json:"TempHighLevel"`
	FreezingLevel string `json:"FreezingLevel"`
}

// GetMountainAreaForecast returns all the mountain areas available from the Metoffice
// in an array.
func (c *Client) GetMountainAreaForecast(area string) (*MountainAreaInput, error) {
	// TODO check that area is a string that's valid

	path := fmt.Sprintf("/public/data/txt/wxfcs/mountainarea/%s/%s", "json", area)
	resp, err := c.Get(path, nil)
	if err != nil {
		return nil, err
	}

	var b *MountainAreaInput
	if err := decodeJSON(&b, resp.Body); err != nil {
		return nil, err
	}

	return b, nil
}

// Location is each of the locations.
type Location struct {
	Id   string `json:"@Id"`
	Name string `json:"@Name"`
}

// Locations indicates the list of locations
type Locations struct {
	Locations []Location `json:"Location"`
}

// MountainAreaSitelist represents the list of mountain areas
// in the metoffice.
type MountainAreaSitelist struct {
	Locations Locations `json:"Locations"`
}

// GetMountainAreaSitelist Returns a list of locations the mountain area
// forecast data feed provides data for .You can use this to find the
// ID of the site that you are interested in.
func (c *Client) GetMountainAreaSitelist() (*MountainAreaSitelist, error) {

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

	return b, nil
}
