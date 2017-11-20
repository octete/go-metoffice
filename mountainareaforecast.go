package metoffice

import (
	"fmt"
	"time"
)

// MountainAreaInput is empty
type MountainAreaInput struct {
	Report Report `json:"report"`
}

// Report is the whole data returned by the MetOffice
type Report struct {
	ID         int64
	FetchData  time.Time
	Location   string     `json:"Location"`
	Issue      string     `json:"Issue"`
	Issued     time.Time  `json:"Issued"`
	Type       string     `json:"Type"`
	ParamUnits ParamUnits `json:"ParamUnits"`
	Evening    Evening    `json:"Evening"`
	DaysInput  DaysInput  `json:"Days"`
}

// DaysInput is
type DaysInput struct {
	DayInput []Day `json:"Day"`
}

// ParamUnits specifies the units used in the report
type ParamUnits struct {
	WindSpeed   string `json:"WindSpeed"`
	MaxGust     string `json:"MaxGust"`
	Temperature string `json:"Temperature"`
	FeelsLike   string `json:"FeelsLike"`
}

// Evening shows the forecast for the evening
// TODO Does this also return something for the Morning?
type Evening struct {
	// XXX TODO check if we can have the validity as a time.Time
	//Validity time.Time `json:"Validity"`
	Validity string `json:"Validity"`
	Summary  string `json:"Summary"`
}

// Day is the first day
type Day struct {
	Validity         string          `json:"Validity"`
	Headline         string          `json:"Headline"`
	Confidence       string          `json:"Confidence"`
	View             string          `json:"View"`
	CloudFreeHillTop string          `json:"CloudFreeHillTop"`
	Weather          string          `json:"Weather"`
	Visibility       string          `json:"Visibility"`
	Hazards          HazardsInput    `json:"Hazards"`
	Periods          PeriodInput     `json:"Periods"`
	GroundConditions GroundCondition `json:"GroundConditions"`
	Temperature      Temperature     `json:"Temperature"`
	Summary          string          `json:"Summary"`
}

// GroundCondition is a summary
type GroundCondition struct {
	Summary string `json:"Summary"` // it normally has a \n at the end
}

// Temperature is a definition of the temperature
type Temperature struct {
	Peak     Peak   `json:"Peak"`
	Valley   Valley `json:"Valley"`
	Freezing string `json:"Freezing"`
}

// Peak shows information on the peaks
type Peak struct {
	Level       string `json:"Level"`
	Description string `json:"$"`
}

// Valley shows information on the valleys.
type Valley struct {
	Title       string `json:"Title"`
	Description string `json:"$"`
}

// HazardsInput is only used to grab all hazzards
type HazardsInput struct {
	HazardsInput []Hazard `json:"Hazard"`
}

// Hazard specify each type of Hazard
type Hazard struct {
	Element    Element    `json:"Element"`
	Likelyhood Likelyhood `json:"Likelyhood"`
}

// Element is each of the Hazard Elements
type Element struct {
	Type        string `json:"Type"`
	Description string `json:"$"`
}

// Likelyhood is each of the hazard likelyhoods
type Likelyhood struct {
	Type        string `json:"Type"`
	Description string `json:"$"`
}

// PeriodInput is the list of periods
type PeriodInput struct {
	PeriodInput []Period `json:"Period"`
}

// Period is the definition of each period
type Period struct {
	Start              string             `json:"Start"`
	End                string             `json:"End"`
	SignificantWeather significantWeather `json:"SignificantWeather"`
	Precipitation      precipitation      `json:"Precipitation"`
	Heights            []Height           `json:"Height"`
	FreezingLevel      string             `json:"FreezingLevel"`
}

type significantWeather struct {
	Code        string `json:"Code"`
	Description string `json:"$"`
}

type precipitation struct {
	Probability string `json:"Probability"`
}

// Height is the different heights
type Height struct {
	Level         string `json:"Level"`
	WindDirection string `json:"WindDirection"`
	WindSpeed     int    `json:"WindSpeed"`
	MaxGust       int    `json:"MaxGust"`
	Temperature   int    `json:"Temperature"`
	FeelsLike     int    `json:"FeelsLike"`
}

// TODO Get a list of all the likelyhoods and element types in an
// ENUM

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
