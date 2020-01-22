package gonube

import "testing"

func TestProducts(t *testing.T) {
	c := getTestClient()
	r, err := c.Products().All(nil)
	if err != nil {
		t.Fatal(err)
	}

	// Debug
	if c.config.Debug {
		t.Logf("%+v", r)
	}

	if len(r) > 0 {
		pr, err := c.Products().Get(r[0].ID, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Debug
		if c.config.Debug {
			t.Logf("%+v", pr)
		}
	}
}
