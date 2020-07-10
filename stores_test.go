package gonube

import "testing"

func TestStores(t *testing.T) {
	c := getTestClient()
	r, err := c.Stores().Get()
	if err != nil {
		t.Fatal(err)
	}

	// Debug
	if c.config.Debug {
		t.Logf("%+v", r)
	}
}
