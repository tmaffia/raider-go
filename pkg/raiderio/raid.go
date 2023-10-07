package raiderio

import (
	"errors"

	"github.com/tmaffia/raiderio/pkg/raiderio/util"
)

// RaidQuery is a struct that represents the query parameters
// sent for a raid request
// Supports optional request fields: difficulty, region, realm, name
type RaidQuery struct {
	Name       string
	Difficulty RaidDifficulty
	Region     string
	Realm      string
	Limit      int
	Page       int
}

// RaidRankings is a struct that represents the response from a
// raid rankings request
type RaidRankings struct {
	RaidRanking []RaidRanking `json:"raidRankings"`
}

// RaidRanking is a struct that represents a raid ranking in a
// raid rankings response from the api
// Unfortunately the "Guild" object differs in structure from the
// guild profile response. This requires a separate struct
type RaidRanking struct {
	Rank         int `json:"rank"`
	RegionalRank int `json:"region_rank"`
	Guild        struct {
		Id      int         `json:"id"`
		Name    string      `json:"name"`
		Faction string      `json:"faction"`
		Realm   util.Realm  `json:"realm"`
		Region  util.Region `json:"region"`
		Path    string      `json:"path"`
		Logo    string      `json:"logo"`
		Color   string      `json:"color"`
	} `json:"guild"`
	EncountersDefeated []struct {
		Slug           string `json:"slug"`
		LastDefeatedAt string `json:"lastDefeated"`
		FirstDefeated  string `json:"firstDefeated"`
	} `json:"encountersDefeated"`
	EncountersPulled []struct {
		Id             int     `json:"id"`
		Slug           string  `json:"slug"`
		Pulls          int     `json:"numPulls"`
		PullsStartedAt string  `json:"pullStartedAt"`
		BestPercent    float32 `json:"bestPercent"`
		IsDefeated     bool    `json:"isDefeated"`
	} `json:"encountersPulled"`
}

// RaidProgression is a struct that contains the raid progression of a guild
// in a guild profile response
type RaidProgression struct {
	Summary     string `json:"summary"`
	Bosses      int    `json:"total_bosses"`
	NormalKills int    `json:"normal_bosses_killed"`
	HeroicKills int    `json:"heroic_bosses_killed"`
	MythicKills int    `json:"mythic_bosses_killed"`
}

// GuildRaidRanking is a struct that contains the raid rankings of a guild
// in a guild profile response
// Includes Normal Heroic and Mythic rankings
type GuildRaidRanking struct {
	Normal struct {
		World  int `json:"world"`
		Region int `json:"region"`
		Realm  int `json:"realm"`
	} `json:"normal"`

	Heroic struct {
		World  int `json:"world"`
		Region int `json:"region"`
		Realm  int `json:"realm"`
	} `json:"heroic"`

	Mythic struct {
		World  int `json:"world"`
		Region int `json:"region"`
		Realm  int `json:"realm"`
	} `json:"mythic"`
}

// Raids is a struct that represents the response from a
// raid static data request
type Raids struct {
	Raids []Raid `json:"raids"`
}

// Raid is a struct that represents a raid in a raid static
// data response. Includes raid encounters and other static data
type Raid struct {
	Id        int    `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Icon      string `json:"icon"`
	Starts    struct {
		Us string `json:"us"`
		Eu string `json:"eu"`
		Tw string `json:"tw"`
		Kr string `json:"kr"`
		Cn string `json:"cn"`
	} `json:"starts"`
	Ends struct {
		Us string `json:"us"`
		Eu string `json:"eu"`
		Tw string `json:"tw"`
		Kr string `json:"kr"`
		Cn string `json:"cn"`
	} `json:"ends"`

	Encounters []Encounter `json:"encounters"`
}

// Encounter is a struct that represents an encounter in a raid
// in a raid static data response
type Encounter struct {
	Id   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// Difficulty is a string type that represents the difficulty of a raid
// in a raid request
type RaidDifficulty string

const (
	Normal RaidDifficulty = "normal"
	Heroic RaidDifficulty = "heroic"
	Mythic RaidDifficulty = "mythic"
)

// validateRaidQuery validates a RaidQuery struct
// ensures that the required parameters are not empty
func validateRaidRankingsQuery(rq *RaidQuery) error {
	if rq.Name == "" {
		return errors.New("no raid name provided")
	}

	if rq.Difficulty == "" {
		return errors.New("no raid difficulty provided")
	}

	if rq.Region == "" {
		return errors.New("no region provided")
	}

	if rq.Limit < 0 {
		return errors.New("limit must be a positive int")
	}

	if rq.Page < 0 {
		return errors.New("page must be a positive int")
	}

	return nil
}
