package metoffice

import "testing"

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
