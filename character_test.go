package raiderio

import (
	"testing"
)

func TestNewCharacterQuery(t *testing.T) {
	cq := CharacterQuery{
		Region: "us",
		Realm:  "illidan",
		Name:   "liquid",
	}

	err := createCharacterQuery(&cq)
	if err != nil {
		t.Errorf("Error creating guild query")
	}
	t.Logf("%+v", cq)
}

func TestNewCharacterQueryWGear(t *testing.T) {
	cq := CharacterQuery{
		Region: "us",
		Realm:  "illidan",
		Name:   "liquid",
		Gear:   true,
	}

	err := createCharacterQuery(&cq)
	if err != nil {
		t.Errorf("Error creating guild query")
	}
	if cq.fields[0] != "gear" {
		t.Errorf("Error creating guild query")
	}
	t.Logf("%+v", cq)
}

func TestNewCharacterQueryWTalents(t *testing.T) {
	cq := CharacterQuery{
		Region:        "us",
		Realm:         "illidan",
		Name:          "liquid",
		TalentLoadout: true,
	}

	err := createCharacterQuery(&cq)
	if err != nil {
		t.Errorf("Error creating guild query")
	}
	if cq.fields[0] != "talent_loadout" {
		t.Errorf("Error creating guild query")
	}
	t.Logf("%+v", cq)
}
