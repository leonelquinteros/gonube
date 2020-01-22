package gonube

import "testing"

func TestOrders(t *testing.T) {
	c := getTestClient()
	r, err := c.Orders().All(nil)
	if err != nil {
		t.Fatal(err)
	}

	// Debug
	if c.config.Debug {
		t.Logf("%+v", r)
	}
}
