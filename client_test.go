package raiderio

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	if c.apiUrl != "https://raider.io/api/v1" {
		t.Errorf("NewClient apiUrl created incorrectly")
	}
}

func TestGetCharacterProfile(t *testing.T) {
	c := NewClient()
	cq, _ := NewCharacterQuery("us", "illidan", "highervalue", nil)

	profile, err := c.GetCharacterProfile(cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetCharacterWGear(t *testing.T) {
	c := NewClient()

	fields := []string{"gear"}
	cq, _ := NewCharacterQuery("us", "illidan", "highervalue", &fields)

	profile, err := c.GetCharacterProfile(cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetCharacterWTalents(t *testing.T) {
	c := NewClient()

	fields := []string{"talents"}
	cq, _ := NewCharacterQuery("us", "illidan",
		"gigavalue", &fields)

	profile, err := c.GetCharacterProfile(cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}
