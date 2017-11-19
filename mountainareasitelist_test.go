package metoffice

import (
	"testing"
)

func TestClient_GetMountainAreaSitelist(t *testing.T) {
	t.Parallel()

	var err error

	var loc *Locations
	record(t, "mountainforecastlist/sitelist", func(c *Client) {
		loc, err = c.GetMountainAreaSitelist()
	})
	if err != nil {
		t.Fatal(err)
	}

	// Check the length of Locations
	length := len(loc.Locations)
	if length <= 0 {
		t.Errorf("GetMountainAreaSitelist: got %d mountain Locations.", length)
	}

	// Check they have IDs and Names
	for _, v := range loc.Locations {
		if len(v.ID) <= 0 {
			t.Errorf("ID %s is empty: %d characters.", v.ID, len(v.ID))
		}
		if len(v.Name) <= 0 {
			t.Errorf("Name %s is empty: %d characters.", v.Name, len(v.Name))
		}
	}
}
