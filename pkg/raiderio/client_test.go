package raiderio

import (
	"testing"

	"github.com/tmaffia/raiderio/pkg/raiderio/util"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	if c.apiUrl != "https://raider.io/api/v1" {
		t.Errorf("NewClient apiUrl created incorrectly")
	}
}

func TestGetCharacterProfile(t *testing.T) {
	c := NewClient()
	cq := CharacterQuery{
		Region: "us",
		Realm:  "illidan",
		Name:   "highervalue",
	}

	profile, err := c.GetCharacter(&cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetCharacterWGear(t *testing.T) {
	c := NewClient()
	cq := CharacterQuery{
		Region: "us",
		Realm:  "illidan",
		Name:   "highervalue",
		Gear:   true,
	}

	profile, err := c.GetCharacter(&cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetCharacterWTalents(t *testing.T) {
	c := NewClient()
	cq := CharacterQuery{
		Region:        "us",
		Realm:         "illidan",
		Name:          "highervalue",
		TalentLoadout: true,
	}

	profile, err := c.GetCharacter(&cq)

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetGuild(t *testing.T) {
	c := NewClient()

	gq := GuildQuery{
		Region: "us",
		Realm:  "illidan",
		Name:   "liquid",
	}

	profile, err := c.GetGuild(&gq)
	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetGuildWMembers(t *testing.T) {
	c := NewClient()

	gq := GuildQuery{
		Region:  "us",
		Realm:   "illidan",
		Name:    "liquid",
		Members: true,
	}

	profile, err := c.GetGuild(&gq)
	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetGuildWRaidProgression(t *testing.T) {
	c := NewClient()
	gq := GuildQuery{
		Region:          "us",
		Realm:           "illidan",
		Name:            "liquid",
		RaidProgression: true,
	}

	profile, err := c.GetGuild(&gq)
	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetGuildWRaidRankings(t *testing.T) {
	c := NewClient()
	gq := GuildQuery{
		Region:       "us",
		Realm:        "illidan",
		Name:         "liquid",
		RaidRankings: true,
	}

	profile, err := c.GetGuild(&gq)
	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetRaids(t *testing.T) {
	c := NewClient()

	raids, err := c.GetRaids(util.Dragonflight)
	if err != nil {
		t.Errorf("Error getting raids")
	}
	t.Logf("%+v", raids)
}

func TestGetRaidRankings(t *testing.T) {
	c := NewClient()
	rq := RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: Mythic,
		Region:     "world",
	}

	rr, err := c.GetRaidRankings(&rq)
	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}
