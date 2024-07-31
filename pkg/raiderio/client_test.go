package raiderio_test

import (
	"testing"

	"github.com/tmaffia/raiderio/pkg/raiderio"
	"github.com/tmaffia/raiderio/pkg/raiderio/expansion"
	"github.com/tmaffia/raiderio/pkg/raiderio/region"
)

func TestNewClient(t *testing.T) {
	c := raiderio.NewClient()

	if c.ApiUrl != "https://raider.io/api/v1" {
		t.Errorf("NewClient apiUrl created incorrectly")
	}
}

func TestGetCharacterProfile(t *testing.T) {
	c := raiderio.NewClient()

	testCases := []struct {
		region         *region.Region
		realm          string
		name           string
		expectedErrMsg string
		expectedName   string
	}{
		{region: region.US, realm: "illidan", name: "highervalue", expectedName: "Highervalue"},
		{region: &region.Region{Slug: "badregion"}, realm: "illidan", name: "impossiblecharactername", expectedErrMsg: "status code: 400, error: bad request, failed to find region badregion"},
		{region: region.US, realm: "illidan", name: "impossiblecharactername", expectedErrMsg: "status code: 400, error: bad request, could not find requested character"},
		{region: region.US, realm: "invalidrealm", name: "highervalue", expectedErrMsg: "status code: 400, error: bad request, failed to find realm invalidrealm in region us"},
	}

	for _, tc := range testCases {
		profile, err := c.GetCharacter(&raiderio.CharacterQuery{
			Region: tc.region,
			Realm:  tc.realm,
			Name:   tc.name,
		})

		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected: %v, got: %v", tc.expectedErrMsg, err.Error())
		}

		if err == nil && profile.Name != tc.expectedName {
			t.Fatalf("expected: %v, got: %v", tc.expectedName, profile.Name)
		}
	}
}

func TestGetCharacterWGear(t *testing.T) {
	profile, err := raiderio.NewClient().GetCharacter(&raiderio.CharacterQuery{
		Region: region.US,
		Realm:  "illidan",
		Name:   "highervalue",
		Gear:   true,
	})

	if err != nil {
		t.Errorf("Error getting character")
	}
	t.Logf("%+v", profile)
}

func TestGetCharacterWTalents(t *testing.T) {
	c := raiderio.NewClient()
	cq := raiderio.CharacterQuery{
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
	t.Logf(profile.Class)
}

func TestGetGuild(t *testing.T) {
	c := raiderio.NewClient()

	gq := raiderio.GuildQuery{
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
	c := raiderio.NewClient()

	profile, err := c.GetGuild((&raiderio.GuildQuery{
		Region:  region.US,
		Realm:   "illidan",
		Name:    "warpath",
		Members: true,
	}))

	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetGuildWRaidProgression(t *testing.T) {
	c := raiderio.NewClient()

	profile, err := c.GetGuild(&raiderio.GuildQuery{
		Region:          region.US,
		Realm:           "illidan",
		Name:            "warpath",
		RaidProgression: true,
	})

	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetGuildWRaidRankings(t *testing.T) {
	c := raiderio.NewClient()
	profile, err := c.GetGuild(&raiderio.GuildQuery{
		Region:       region.US,
		Realm:        "illidan",
		Name:         "warpath",
		RaidRankings: true,
	})

	if err != nil {
		t.Errorf("Error getting guild")
	}
	t.Logf("%+v", profile)
}

func TestGetRaids(t *testing.T) {
	c := raiderio.NewClient()
	raids, err := c.GetRaids(expansion.Dragonflight)
	if err != nil {
		t.Errorf("Error getting raids")
	}
	t.Logf("%+v", raids)
}

func TestGetRaidRankings(t *testing.T) {
	c := raiderio.NewClient()

	rr, err := c.GetRaidRankings(&raiderio.RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: raiderio.MythicRaid,
		Region:     region.WORLD,
	})

	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}

func TestGetRaidRankingsWRealm(t *testing.T) {
	c := raiderio.NewClient()

	rr, err := c.GetRaidRankings(&raiderio.RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: raiderio.MythicRaid,
		Region:     region.US,
		Realm:      "illidan",
	})

	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}

func TestGetRaidRankingsWLimit(t *testing.T) {
	c := raiderio.NewClient()

	rr, err := c.GetRaidRankings(&raiderio.RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: raiderio.MythicRaid,
		Region:     region.US,
		Limit:      2,
	})

	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
	t.Logf("%+v", rr)
}
