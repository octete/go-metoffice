package metoffice

import "testing"

func TestClient_GetMountainForecast(t *testing.T) {
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

}
