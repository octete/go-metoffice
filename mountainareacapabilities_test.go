package metoffice

import (
	"testing"
	"time"
)

func TestClient_GetMountainAreaCapabilities(t *testing.T) {
	t.Parallel()

	var err error

	// Get
	var mfItems []MountainForecastItem
	record(t, "mountainforecastlist/get", func(c *Client) {
		mfItems, err = c.GetMountainAreaCapabilities()
	})
	if err != nil {
		t.Fatal(err)
	}

	/*
		type MountainForecastItem struct {
			DataDate   time.Time `json:"DataDate"`
			ValidFrom  time.Time `json:"ValidFrom"`
			ValidTo    time.Time `json:"ValidTo"`
			CreateDate time.Time `json:"CreateDate"`
			URI        string    `json:"URI"`
			Area       string    `json:"Area"`
			Risk       string    `json:"Risk"`
		}
	*/
	for _, v := range mfItems {
		if len(v.URI) <= 0 {
			t.Errorf("URI does not contain anything: %s", v.URI)
		}
		if len(v.Area) <= 0 {
			t.Errorf("Area does not contain anything: %s", v.URI)
		}
		if len(v.Risk) <= 0 {
			t.Errorf("Risk does not contain anything: %s", v.URI)
		}

	}

	// TODO Improve this
	for _, v := range mfItems {

		if len(v.DataDate.Format(time.UnixDate)) <= 0 {
			t.Errorf("DataDate does not have a valid month: %s", v.DataDate.Format(time.UnixDate))
		}
		if len(v.ValidFrom.Format(time.UnixDate)) <= 0 {
			t.Errorf("DataDate does not have a valid month: %s", v.ValidFrom.Format(time.UnixDate))
		}
		if len(v.ValidTo.Format(time.UnixDate)) <= 0 {
			t.Errorf("DataDate does not have a valid month: %s", v.ValidTo.Format(time.UnixDate))
		}
		if len(v.CreateDate.Format(time.UnixDate)) <= 0 {
			t.Errorf("DataDate does not have a valid month: %s", v.CreateDate.Format(time.UnixDate))
		}
	}
}
