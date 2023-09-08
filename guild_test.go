package raiderio

import "testing"

func TestNewGuildQuery(t *testing.T) {
	name := "warpath"
	region := "us"
	realm := "illidan"
	gq, err := NewGuildQuery(region, realm, name, nil)

	if err != nil {
		t.Errorf("Error creating guild query")
	}
	t.Logf("%+v", gq)
}
