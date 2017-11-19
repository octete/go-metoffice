package metoffice

import (
	"testing"
)

func TestClient_MountainForecast(t *testing.T) {
	t.Parallel()

	var err error

	// Get
	var ns []MountainForecastItem
	record(t, "mountainforecastlist/get", func(c *Client) {
		ns, err = c.ListMountainForecasts()
	})
	if err != nil {
		t.Fatal(err)
	}
	/*
	   if s.Name != ns.Name {
	       t.Errorf("bad name: %q (%q)", s.Name, ns.Name)
	   }
	   if s.Comment != ns.Comment {
	       t.Errorf("bad comment: %q (%q)", s.Comment, ns.Comment)
	   }

	   if ns.CreatedAt == "" {
	       t.Errorf("Bad created at: empty")
	   }

	   if ns.UpdatedAt == "" {
	       t.Errorf("Bad updated at: empty")
	   }

	   if ns.DeletedAt != "" {
	       t.Errorf("Bad deleted at: %s", ns.DeletedAt)
	   }
	*/
}

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

	// Check they have IDs and Names
	for _, v := range loc.Locations {
		if len(v.ID) <= 0 {
			t.Errorf("ID %s is empty: %d characters.", v.ID, len(v.ID))
		}
		if len(v.Name) <= 0 {
			t.Errorf("Name %s is empty: %d characters.", v.Name, len(v.Name))
		}
	}

	// Check the length
	length := len(loc.Locations)
	if length <= 0 {
		t.Errorf("GetMountainAreaSitelist: got %d mountain Locations.", length)
	}

}
