package raiderio

import (
	"encoding/json"
	"errors"

	"github.com/tmaffia/raiderio/pkg/raiderio/region"
)

// GuildQuery is a struct that represents the query parameters
// sent for a guild profile request
// Supports optional request fields: members, raid_progression, raid_rankings
type GuildQuery struct {
	Region          *region.Region
	Realm           string
	Name            string
	Members         bool
	RaidProgression bool
	RaidRankings    bool
	fields          []string
}

// Guild is a struct that represents the response from
// a guild profile request
type Guild struct {
	Name            string                      `json:"name"`
	Faction         string                      `json:"faction"`
	Region          string                      `json:"region"`
	Realm           string                      `json:"realm"`
	LastCrawledAt   string                      `json:"last_crawled_at"`
	ProfileUrl      string                      `json:"profile_url"`
	Members         []Member                    `json:"members"`
	RaidProgression GuildRaidProgression        `json:"raid_progression"`
	RaidRankings    map[string]GuildRaidRanking `json:"raid_rankings"`
}

// Member is a struct that represents a member of a guild
// in a guild profile response
type Member struct {
	Rank      int       `json:"rank"`
	Character Character `json:"character"`
}

// RaidProgression is a struct that contains the raid progression of a guild
// in a guild profile response
// Currently supports Dragonflight raids
type GuildRaidProgression struct {
	Amirdrassil          RaidProgression `json:"amirdrassil-amirdrassil-the-dreams-hope"`
	Aberrus              RaidProgression `json:"aberrus-the-shadowed-crucible"`
	VaultOfTheIncarnates RaidProgression `json:"vault-of-the-incarnates"`
}

// createGuildQuery creates and validates a GuildQuery struct
// It returns an error if any of the required parameters are empty
// or if the fields are invalid
func createGuildQuery(gq *GuildQuery) error {
	if gq.Region == nil {
		return ErrInvalidRegion
	}

	if gq.Realm == "" {
		return ErrInvalidRealm
	}

	if gq.Name == "" {
		return ErrInvalidGuildName
	}

	if gq.Members {
		gq.fields = append(gq.fields, "members")
	}

	if gq.RaidProgression {
		gq.fields = append(gq.fields, "raid_progression")
	}

	if gq.RaidRankings {
		gq.fields = append(gq.fields, "raid_rankings")
	}
	return nil
}

func (g *Guild) GetGuildRaidRankBySlug(slug string) (*GuildRaidRanking, error) {
	if g.RaidRankings == nil {
		return nil, errors.New("guild raid rankings " + ErrFieldMissing.Error())
	}

	gr, ok := g.RaidRankings[slug]
	if !ok {
		return nil, ErrInvalidRaid
	}

	return &gr, nil
}

func unmarshalGuild(body []byte) (*Guild, error) {
	var profile Guild
	err := json.Unmarshal(body, &profile)
	if err != nil {
		return nil, errors.New("error unmarshalling guild profile")
	}

	for k := range profile.RaidRankings {
		if entry, ok := profile.RaidRankings[k]; ok {
			entry.RaidSlug = k
			profile.RaidRankings[k] = entry
		}
	}
	return &profile, nil
}
