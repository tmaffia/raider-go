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
		{region: region.US, realm: "", name: "highervalue", expectedErrMsg: "invalid realm"},
		{region: region.US, realm: "illidan", name: "", expectedErrMsg: "invalid character name"},
		{region: nil, realm: "illidan", name: "highervalue", expectedErrMsg: "invalid region"},
		{region: &region.Region{Slug: "badregion"}, realm: "illidan", name: "impossiblecharactername", expectedErrMsg: "invalid region"},
		{region: region.US, realm: "illidan", name: "impossiblecharactername", expectedErrMsg: "character not found"},
		{region: region.US, realm: "invalidrealm", name: "highervalue", expectedErrMsg: "invalid realm"},
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
			t.Fatalf("character name expected: %v, got: %v", tc.expectedName, profile.Name)
		}
	}
}

func TestGetCharacterWGear(t *testing.T) {
	c := raiderio.NewClient()

	testCases := []struct {
		region         *region.Region
		realm          string
		name           string
		expectedErrMsg string
		expectedName   string
	}{
		{region: region.US, realm: "illidan", name: "highervalue", expectedName: "Highervalue"},
	}

	for _, tc := range testCases {
		profile, err := c.GetCharacter(&raiderio.CharacterQuery{
			Region: tc.region,
			Realm:  tc.realm,
			Name:   tc.name,
			Gear:   true,
		})

		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected: %v, got: %v", tc.expectedErrMsg, err.Error())
		}

		if err == nil && profile.Name != tc.expectedName {
			t.Fatalf("character name expected: %v, got: %v. item level equipped: %d", tc.expectedName, profile.Name, profile.Gear.ItemLevelEquipped)
		}

		if err == nil && !(profile.Gear.ItemLevelEquipped > 0) {
			t.Fatalf("character item level equipped: %d, expected > 0", profile.Gear.ItemLevelEquipped)
		}
	}
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
	if err == nil && profile.TalentLoadout.LoadoutText == "" {
		t.Fatalf("talent loadout: %v expected to not be empty", profile.TalentLoadout.LoadoutText)
	}
}

func TestGetGuild(t *testing.T) {
	c := raiderio.NewClient()

	testCases := []struct {
		region         *region.Region
		realm          string
		name           string
		expectedErrMsg string
		expectedName   string
	}{
		{region: region.US, realm: "illidan", name: "warpath", expectedName: "Warpath"},
		{region: region.US, realm: "", name: "warpath", expectedErrMsg: "invalid realm"},
		{region: region.US, realm: "illidan", name: "", expectedErrMsg: "invalid guild name"},
		{region: nil, realm: "illidan", name: "highervalue", expectedErrMsg: "invalid region"},
		{region: &region.Region{Slug: "badregion"}, realm: "illidan", name: "warpath", expectedErrMsg: "invalid region"},
		{region: region.US, realm: "illidan", name: "impossible_guild_name", expectedErrMsg: "guild not found"},
		{region: region.US, realm: "invalidrealm", name: "highervalue", expectedErrMsg: "invalid realm"},
	}

	for _, tc := range testCases {
		profile, err := c.GetGuild(&raiderio.GuildQuery{
			Region: tc.region,
			Realm:  tc.realm,
			Name:   tc.name,
		})

		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected: %v, got: %v", tc.expectedErrMsg, err.Error())
		}

		if err == nil && profile.Name != tc.expectedName {
			t.Fatalf("guild name expected: %v, got: %v.", tc.expectedName, profile.Name)
		}
	}
}

func TestGetGuildWMembers(t *testing.T) {
	c := raiderio.NewClient()
	testCases := []struct {
		region *region.Region
		realm  string
		name   string
	}{
		{region: region.US, realm: "illidan", name: "warpath"},
	}

	for range testCases {
		profile, err := c.GetGuild((&raiderio.GuildQuery{
			Region:  region.US,
			Realm:   "illidan",
			Name:    "warpath",
			Members: true,
		}))

		if err != nil {
			t.Fatalf("Error getting guild")
		}

		if !(len(profile.Members) > 0) {
			t.Fatalf("Error getting guild members")
		}
	}

}

func TestGetGuildWRaidProgression(t *testing.T) {
	c := raiderio.NewClient()
	testCases := []struct {
		region *region.Region
		realm  string
		name   string
	}{
		{region: region.US, realm: "illidan", name: "warpath"},
	}

	for range testCases {
		profile, err := c.GetGuild(&raiderio.GuildQuery{
			Region:          region.US,
			Realm:           "illidan",
			Name:            "warpath",
			RaidProgression: true,
		})

		if err != nil {
			t.Errorf("Error getting guild")
		}

		if profile.RaidProgression.Aberrus.Summary == "" {
			t.Fatalf("Error getting guild raid progression, got: %v", profile.RaidProgression.Amirdrassil.Summary)
		}
	}
}

func TestGetGuildWRaidRankings(t *testing.T) {
	c := raiderio.NewClient()
	testCases := []struct {
		region       *region.Region
		realm        string
		name         string
		raidName     string
		expectedRank int
	}{
		{region: region.US, realm: "illidan", name: "warpath",
			raidName: "aberrus-the-shadowed-crucible", expectedRank: 158},
	}

	for _, tc := range testCases {
		profile, err := c.GetGuild(&raiderio.GuildQuery{
			Region:       region.US,
			Realm:        "illidan",
			Name:         "warpath",
			RaidRankings: true,
		})

		if err != nil {
			t.Errorf("Error getting guild")
		}

		rank := profile.RaidRankings[tc.raidName]

		if rank.Mythic.World != tc.expectedRank {
			t.Fatalf("mythic guild ranking for raid: %v, got: %d, expected: %d",
				rank.RaidSlug, rank.Mythic.World, tc.expectedRank)
		}
	}
}

func TestGetRaids(t *testing.T) {
	c := raiderio.NewClient()
	testCases := []struct {
		expansion        expansion.Expansion
		raidName         string
		expectedRaidName string
		expectedErrMsg   string
	}{
		{expansion: expansion.Dragonflight},
		{expansion: 2, expectedErrMsg: "unexpected error"},
	}

	for _, tc := range testCases {
		_, err := c.GetRaids(tc.expansion)

		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected error: %v, got %v", tc.expectedErrMsg, err.Error())
		}
	}

	_, err := c.GetRaids(expansion.Dragonflight)
	if err != nil {
		t.Errorf("Error getting raids")
	}
}

func TestGetRaidRankings(t *testing.T) {
	c := raiderio.NewClient()

	_, err := c.GetRaidRankings(&raiderio.RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: raiderio.MythicRaid,
		Region:     region.WORLD,
	})

	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
}

func TestGetRaidRankingsWRealm(t *testing.T) {
	c := raiderio.NewClient()

	_, err := c.GetRaidRankings(&raiderio.RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: raiderio.MythicRaid,
		Region:     region.US,
		Realm:      "illidan",
	})

	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
}

func TestGetRaidRankingsWLimit(t *testing.T) {
	c := raiderio.NewClient()

	_, err := c.GetRaidRankings(&raiderio.RaidQuery{
		Name:       "aberrus-the-shadowed-crucible",
		Difficulty: raiderio.MythicRaid,
		Region:     region.US,
		Limit:      2,
	})

	if err != nil {
		t.Errorf("Error getting raid rankings: " + err.Error())
	}
}
