package raiderio_test

import (
	"testing"

	"github.com/tmaffia/raiderio/pkg/raiderio"
	"github.com/tmaffia/raiderio/pkg/raiderio/region"
)

func TestGetGuildRaidRankBySlug(t *testing.T) {
	c := raiderio.NewClient()

	testCases := []struct {
		region              *region.Region
		realm               string
		name                string
		includeRandRankings bool
		raidSlug            string
		expectedRaidRank    int
		expectedErrMsg      string
	}{
		{region: region.US, realm: "illidan", name: "warpath", raidSlug: "aberrus-the-shadowed-crucible", expectedRaidRank: 158, includeRandRankings: true},
		{region: region.US, realm: "illidan", name: "warpath", raidSlug: "invalid raid slug", expectedErrMsg: "invalid raid", includeRandRankings: true},
		{region: region.US, realm: "illidan", name: "warpath", raidSlug: "aberrus-the-shadowed-crucible",
			expectedErrMsg: "guild raid rankings field missing from api response", includeRandRankings: false},
	}

	for _, tc := range testCases {
		profile, _ := c.GetGuild(&raiderio.GuildQuery{
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
