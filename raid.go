package raiderio

type Raid struct {
	Name       string `json:"name"`
	Faction    string `json:"faction"`
	Region     string `json:"region"`
	Realm      string `json:"realm"`
	ProfileUrl string `json:"profile_url"`
}
