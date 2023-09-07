package raiderio

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient()

	if err != nil {
		t.Errorf("NewClient creation fail")
	}
	if c.apiUrl != "https://raider.io/api/v1" {
		t.Errorf("NewClient apiUrl created incorrectly")
	}
}

func TestGetCharacterProfile(t *testing.T) {
	c, _ := NewClient()
	cq, _ := NewCharacterQuery("us", "illidan", "highervalue", nil)

	profile, err := c.GetCharacterProfile(cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetCharacterWithGear(t *testing.T) {
	c, _ := NewClient()

	fields := []string{"gear"}
	cq, _ := NewCharacterQuery("us", "illidan", "highervalue", &fields)

	profile, err := c.GetCharacterProfile(cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}
