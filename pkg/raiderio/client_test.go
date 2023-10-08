package raiderio

import (
	"testing"

	"github.com/tmaffia/raiderio/pkg/raiderio/expansion"
	"github.com/tmaffia/raiderio/pkg/raiderio/region"
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
		Region: region.US,
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
		Region: region.US,
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
		Region:        region.US,
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
		Region: region.US,
		Realm:  "illidan",
		Name:   "warpath",
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
		Region:  region.US,
		Realm:   "illidan",
		Name:    "warpath",
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
		Region:          region.US,
		Realm:           "illidan",
		Name:            "warpath",
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
		Region:       region.US,
		Realm:        "illidan",
		Name:         "warpath",
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

	raids, err := c.GetRaids(expansion.Dragonflight)
	if err != nil {
		t.Errorf("Error getting raids")
	}
	t.Logf("%+v", raids)
}

func TestGetRaidRankings(t *testing.T) {
	c := NewClient()
	rq := RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: MythicRaid,
		Region:     region.WORLD,
	}

	rr, err := c.GetRaidRankings(&rq)
	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}

func TestGetRaidRankingsWRealm(t *testing.T) {
	c := NewClient()
	rq := RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: MythicRaid,
		Region:     region.US,
		Realm:      "illidan",
	}

	rr, err := c.GetRaidRankings(&rq)
	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}

func TestGetRaidRankingsWLimit(t *testing.T) {
	c := NewClient()
	rq := RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: MythicRaid,
		Region:     region.US,
		Limit:      2,
	}

	rr, err := c.GetRaidRankings(&rq)
	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}
