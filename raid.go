package raiderio

// RaidProgression is a struct that contains the raid progression of a guild
// in a guild profile response
type RaidProgression struct {
	Summary     string `json:"summary"`
	Bosses      int    `json:"total_bosses"`
	NormalKills int    `json:"normal_bosses_killed"`
	HeroicKills int    `json:"heroic_bosses_killed"`
	MythicKills int    `json:"mythic_bosses_killed"`
}

// RaidRanking is a struct that contains the raid rankings of a guild
// in a guild profile response
// Includes Normal Heroic and Mythic rankings
type RaidRanking struct {
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
