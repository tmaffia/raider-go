package raiderio_test

import (
	"testing"

	"github.com/tmaffia/raiderio"
	"github.com/tmaffia/raiderio/regions"
)

func TestGetGuildRaidRankBySlug(t *testing.T) {
	testCases := []struct {
		region              *regions.Region
		realm               string
		name                string
		includeRandRankings bool
		raidSlug            string
		expectedRaidRank    int
		expectedErrMsg      string
	}{
		{region: regions.US, realm: "illidan", name: "warpath", raidSlug: "nerubar-palace", expectedRaidRank: 92, includeRandRankings: true},
		{region: regions.US, realm: "illidan", name: "warpath", raidSlug: "invalid raid slug", expectedErrMsg: "invalid raid", includeRandRankings: true},
		{region: regions.US, realm: "illidan", name: "warpath", raidSlug: "nerubar-palace",
			expectedErrMsg: "guild raid rankings field missing from api response", includeRandRankings: false},
	}

	for _, tc := range testCases {
		profile, _ := c.GetGuild(defaultCtx, &raiderio.GuildQuery{
			Region:       tc.region,
			Realm:        tc.realm,
			Name:         tc.name,
			RaidRankings: tc.includeRandRankings,
		})

		rank, err := profile.GetGuildRaidRankBySlug(tc.raidSlug)
		if err != nil && err.Error() != tc.expectedErrMsg {
			t.Fatalf("expected error: %v, got: %v", tc.expectedErrMsg, err.Error())
		}

		if err == nil && rank.Mythic.World != tc.expectedRaidRank {
			t.Fatalf("mythic guild ranking for raid: %v, got: %d, expected: %d",
				rank.RaidSlug, rank.Mythic.World, tc.expectedRaidRank)
		}
	}
}
