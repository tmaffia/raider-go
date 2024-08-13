package raiderio_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/tmaffia/raiderio"
)

var c *raiderio.Client
var defaultCtx context.Context

func setup() {
	c = raiderio.NewClient()
	defaultCtx = context.Background()
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNewClient(t *testing.T) {
	if c.ApiUrl != "https://raider.io/api/v1" {
		t.Errorf("NewClient apiUrl created incorrectly")
	}
}

func TestGetCharacterProfile(t *testing.T) {
	testCases := []struct {
		timeout        bool
		region         *raiderio.Region
		realm          string
		name           string
		expectedErrMsg string
		expectedName   string
	}{
		{region: raiderio.Regions.US, realm: "illidan", name: "highervalue", expectedName: "Highervalue"},
		{region: raiderio.Regions.US, realm: "", name: "highervalue", expectedErrMsg: "invalid realm"},
		{region: raiderio.Regions.US, realm: "illidan", name: "", expectedErrMsg: "invalid character name"},
		{region: nil, realm: "illidan", name: "highervalue", expectedErrMsg: "invalid region"},
		{region: &raiderio.Region{Slug: "badregion"}, realm: "illidan", name: "impossiblecharactername", expectedErrMsg: "invalid region"},
		{region: raiderio.Regions.US, realm: "illidan", name: "impossiblecharactername", expectedErrMsg: "character not found"},
		{region: raiderio.Regions.US, realm: "invalidrealm", name: "highervalue", expectedErrMsg: "invalid realm"},
		{timeout: true, region: raiderio.Regions.US, realm: "illidan", name: "highervalue", expectedErrMsg: "raiderio api request timeout"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(defaultCtx, time.Millisecond*1)
			defer cancel()
		}

		profile, err := c.GetCharacter(ctx, &raiderio.CharacterQuery{
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
	testCases := []struct {
		timeout        bool
		region         *raiderio.Region
		realm          string
		name           string
		expectedErrMsg string
		expectedName   string
	}{
		{region: raiderio.Regions.US, realm: "illidan", name: "highervalue", expectedName: "Highervalue"},
		{timeout: true, region: raiderio.Regions.US, realm: "illidan", name: "highervalue", expectedErrMsg: "raiderio api request timeout"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(ctx, time.Millisecond*1)
			defer cancel()
		}

		profile, err := c.GetCharacter(ctx, &raiderio.CharacterQuery{
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
	cq := raiderio.CharacterQuery{
		Region:        raiderio.Regions.US,
		Realm:         "illidan",
		Name:          "highervalue",
		TalentLoadout: true,
	}

	profile, err := c.GetCharacter(defaultCtx, &cq)
	if err == nil && profile.TalentLoadout.LoadoutText == "" {
		t.Fatalf("talent loadout: %v expected to not be empty", profile.TalentLoadout.LoadoutText)
	}
}

func TestGetGuild(t *testing.T) {
	testCases := []struct {
		timeout        bool
		region         *raiderio.Region
		realm          string
		name           string
		expectedErrMsg string
		expectedName   string
	}{
		{region: raiderio.Regions.US, realm: "illidan", name: "warpath", expectedName: "Warpath"},
		{region: raiderio.Regions.US, realm: "", name: "warpath", expectedErrMsg: "invalid realm"},
		{region: raiderio.Regions.US, realm: "illidan", name: "", expectedErrMsg: "invalid guild name"},
		{region: nil, realm: "illidan", name: "highervalue", expectedErrMsg: "invalid region"},
		{region: &raiderio.Region{Slug: "badregion"}, realm: "illidan", name: "warpath", expectedErrMsg: "invalid region"},
		{region: raiderio.Regions.US, realm: "illidan", name: "impossible_guild_name", expectedErrMsg: "guild not found"},
		{region: raiderio.Regions.US, realm: "invalidrealm", name: "highervalue", expectedErrMsg: "invalid realm"},
		{timeout: true, region: raiderio.Regions.US, realm: "illidan", name: "highervalue", expectedErrMsg: "raiderio api request timeout"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(ctx, time.Millisecond*1)
			defer cancel()
		}

		profile, err := c.GetGuild(ctx, &raiderio.GuildQuery{
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
	testCases := []struct {
		region *raiderio.Region
		realm  string
		name   string
	}{
		{region: raiderio.Regions.US, realm: "illidan", name: "warpath"},
	}

	for range testCases {
		profile, err := c.GetGuild(defaultCtx, &raiderio.GuildQuery{
			Region:  raiderio.Regions.US,
			Realm:   "illidan",
			Name:    "warpath",
			Members: true,
		})

		if err != nil {
			t.Fatalf("Error getting guild")
		}

		if !(len(profile.Members) > 0) {
			t.Fatalf("Error getting guild members")
		}
	}

}

func TestGetGuildWRaidProgression(t *testing.T) {
	testCases := []struct {
		region *raiderio.Region
		realm  string
		name   string
	}{
		{region: raiderio.Regions.US, realm: "illidan", name: "warpath"},
	}

	for range testCases {
		profile, err := c.GetGuild(defaultCtx, &raiderio.GuildQuery{
			Region:          raiderio.Regions.US,
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
	testCases := []struct {
		timeout        bool
		region         *raiderio.Region
		realm          string
		name           string
		raidName       string
		expectedRank   int
		expectedErrMsg string
	}{
		{region: raiderio.Regions.US, realm: "illidan", name: "warpath",
			raidName: "aberrus-the-shadowed-crucible", expectedRank: 158},
		{timeout: true, region: raiderio.Regions.US, realm: "illidan", name: "warpath",
			raidName:       "aberrus-the-shadowed-crucible",
			expectedErrMsg: "raiderio api request timeout"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(ctx, time.Millisecond*1)
			defer cancel()
		}

		profile, err := c.GetGuild(ctx, &raiderio.GuildQuery{
			Region:       raiderio.Regions.US,
			Realm:        "illidan",
			Name:         "warpath",
			RaidRankings: true,
		})

		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("error message got: %v, expected: %v",
				err.Error(), tc.expectedErrMsg)
		}

		if err == nil {
			rank := profile.RaidRankings[tc.raidName]
			if rank.Mythic.World != tc.expectedRank {
				t.Fatalf("mythic guild ranking for raid: %v, got: %d, expected: %d",
					rank.RaidSlug, rank.Mythic.World, tc.expectedRank)
			}
		}
	}
}

func TestGetGuildBossKill(t *testing.T) {
	testCases := []struct {
		region                *raiderio.Region
		realm                 string
		guildName             string
		raidSlug              string
		bossSlug              string
		difficulty            raiderio.RaidDifficulty
		expectedDefeatedAt    string
		expectedCharacterName string
		expectedErrMsg        string
		timeout               bool
	}{
		{region: raiderio.Regions.US, realm: "illidan", guildName: "warpath",
			raidSlug: "vault-of-the-incarnates", bossSlug: "terros",
			difficulty: raiderio.Difficulty.MythicRaid, expectedCharacterName: "Drbananaphd"},
		{region: nil, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			guildName: "warpath", raidSlug: "vault-of-the-incarnates",
			bossSlug: "terros", expectedErrMsg: "invalid region"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid,
			realm: "invalid-realm", guildName: "warpath", raidSlug: "vault-of-the-incarnates",
			bossSlug: "terros", expectedErrMsg: "invalid realm"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid,
			guildName: "warpath", raidSlug: "vault-of-the-incarnates",
			bossSlug: "terros", expectedErrMsg: "invalid realm"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			guildName: "impossible-guild_name", raidSlug: "vault-of-the-incarnates",
			bossSlug: "terros", expectedErrMsg: "guild not found"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			raidSlug: "vault-of-the-incarnates", bossSlug: "terros",
			expectedErrMsg: "invalid guild name"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			guildName: "warpath", raidSlug: "invalid-raid-slug", bossSlug: "terros",
			expectedErrMsg: "invalid raid"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			guildName: "warpath", bossSlug: "terros",
			expectedErrMsg: "invalid raid name"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			guildName: "warpath", raidSlug: "vault-of-the-incarnates",
			bossSlug: "invalid-boss-slug", expectedErrMsg: "invalid boss"},
		{region: raiderio.Regions.US, difficulty: raiderio.Difficulty.MythicRaid, realm: "illidan",
			guildName: "warpath", raidSlug: "vault-of-the-incarnates",
			expectedErrMsg: "invalid boss"},
		{region: raiderio.Regions.US, realm: "illidan", guildName: "warpath",
			raidSlug: "vault-of-the-incarnates", bossSlug: "terros",
			expectedErrMsg: "invalid raid difficulty"},
		{timeout: true, region: raiderio.Regions.US, realm: "illidan", guildName: "warpath",
			raidSlug: "vault-of-the-incarnates", bossSlug: "terros",
			difficulty:     raiderio.Difficulty.MythicRaid,
			expectedErrMsg: "raiderio api request timeout"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(ctx, time.Millisecond*1)
			defer cancel()
		}

		k, err := c.GetGuildBossKill(ctx, &raiderio.GuildBossKillQuery{
			Region:     tc.region,
			Realm:      tc.realm,
			GuildName:  tc.guildName,
			RaidSlug:   tc.raidSlug,
			BossSlug:   tc.bossSlug,
			Difficulty: tc.difficulty,
		})

		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("error message got: %v, expected: %v", err.Error(), tc.expectedErrMsg)
		}

		if err == nil && !killIncludesCharacter(k, tc.expectedCharacterName) {
			t.Fatalf("boss kill character name expected: %v", tc.expectedCharacterName)
		}
	}
}

func TestGetRaids(t *testing.T) {
	testCases := []struct {
		timeout          bool
		expansion        raiderio.Expansion
		raidName         string
		expectedRaidName string
		expectedErrMsg   string
	}{
		{expansion: raiderio.Expansions.Dragonflight, raidName: "aberrus-the-shadowed-crucible", expectedRaidName: "Aberrus, the Shadowed Crucible"},
		{timeout: true, expansion: raiderio.Expansions.Dragonflight, raidName: "aberrus-the-shadowed-crucible", expectedErrMsg: "raiderio api request timeout"},
		{expansion: 2, expectedErrMsg: "unsupported expansion"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(ctx, time.Millisecond*1)
			defer cancel()
		}

		raids, err := c.GetRaids(ctx, tc.expansion)
		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected error: %v, got %v", tc.expectedErrMsg, err.Error())
		}

		if err == nil {
			raid, _ := raids.GetRaidBySlug(tc.raidName)
			if raid.Name != tc.expectedRaidName {
				t.Fatalf("expected raid name: %v, got: %v", tc.expectedRaidName, raid.Name)
			}
		}

	}
}

func TestGetRaidRankings(t *testing.T) {
	testCases := []struct {
		timeout                bool
		slug                   string
		difficulty             raiderio.RaidDifficulty
		region                 *raiderio.Region
		realm                  string
		limit                  int
		page                   int
		expectedErrMsg         string
		expectedRank1GuildName string
	}{
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.WORLD,
			expectedRank1GuildName: "Liquid"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US,
			realm: "proudmoore", expectedRank1GuildName: "The Royal Knights"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: "mythic", region: raiderio.Regions.EU, expectedRank1GuildName: "Echo"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US,
			realm: "illidan", expectedRank1GuildName: "Liquid"},
		{slug: "invalid raid slug", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US, realm: "illidan",
			expectedErrMsg: "invalid raid"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: "mythic", region: nil, realm: "illidan", expectedErrMsg: "invalid region"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: "", region: raiderio.Regions.US, realm: "illidan",
			expectedErrMsg: "invalid raid difficulty"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: "invalid-difficulty", region: raiderio.Regions.US, realm: "illidan",
			expectedErrMsg: "invalid raid difficulty"},
		{slug: "", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US, realm: "illidan",
			expectedErrMsg: "invalid raid name"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.WORLD,
			expectedRank1GuildName: "Liquid", limit: 20},
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.WORLD, limit: -20,
			expectedErrMsg: "limit must be a positive int"},
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US,
			expectedRank1GuildName: "Accession", limit: 40, page: 2},
		{slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US, limit: 40,
			page: -2, expectedErrMsg: "page must be a positive int"},
		{timeout: true, slug: "aberrus-the-shadowed-crucible", difficulty: raiderio.Difficulty.MythicRaid, region: raiderio.Regions.US,
			expectedErrMsg: "raiderio api request timeout"},
	}

	for _, tc := range testCases {
		ctx := defaultCtx
		var cancel context.CancelFunc
		if tc.timeout {
			ctx, cancel = context.WithTimeout(ctx, time.Millisecond*1)
			defer cancel()
		}

		rankings, err := c.GetRaidRankings(ctx, &raiderio.RaidQuery{
			Slug:       tc.slug,
			Difficulty: raiderio.RaidDifficulty(tc.difficulty),
			Region:     tc.region,
			Realm:      tc.realm,
			Limit:      tc.limit,
			Page:       tc.page,
		})
		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected error: %v, got: %v", tc.expectedErrMsg, err.Error())
		}

		if err == nil && rankings.RaidRanking[0].Guild.Name != tc.expectedRank1GuildName {
			t.Fatalf("expected guild name: %v, got: %v", tc.expectedRank1GuildName, rankings.RaidRanking[0].Guild.Name)
		}

		if err == nil && tc.limit != 0 {
			if len(rankings.RaidRanking) != tc.limit {
				t.Fatalf("expected results limit: %v, got: %v", tc.limit, len(rankings.RaidRanking))
			}

		}
	}
}

// Tests if character is a part of the particular boss kill
func killIncludesCharacter(k *raiderio.BossKill, c string) bool {
	for _, v := range k.Roster {
		if v.Name == c {
			return true
		}
	}
	return false
}
