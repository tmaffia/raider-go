package raiderio

import "errors"

type CharacterQuery struct {
	Region string   `json:"region"`
	Realm  string   `json:"realm"`
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
}

type CharacterProfile struct {
	Name              string `json:"name"`
	Race              string `json:"race"`
	Class             string `json:"class"`
	ActiveSpec        string `json:"active_spec_name"`
	ActiveRole        string `json:"active_spec_role"`
	Gender            string `json:"gender"`
	Faction           string `json:"faction"`
	AchievementPoints int64  `json:"achievement_points"`
	HonorableKills    int64  `json:"honorable_kills"`
	ThumbnailUrl      string `json:"thumbnail_url"`
	Region            string `json:"region"`
	Realm             string `json:"realm"`
	Gear              Gear   `json:"gear"`
	ProfileUrl        string `json:"profile_url"`
	ProfileBanner     string `json:"profile_banner"`
}

type Gear struct {
	ItemLevelEquipped int `json:"item_level_equipped"`
	ItemLevelTotal    int `json:"item_level_total"`
	// Items             []Item `json:"items"`
}

type Item struct {
}

func NewCharacterQuery(
	region string,
	realm string,
	name string,
	fields *[]string) (*CharacterQuery, error) {

	if region == "" {
		return nil, errors.New("region error")
	}

	if realm == "" {
		return nil, errors.New("realm error")
	}

	if name == "" {
		return nil, errors.New("name error")
	}

	cq := CharacterQuery{
		Region: region,
		Realm:  realm,
		Name:   name,
	}

	if fields != nil {
		cq.Fields = *fields
	}

	return &cq, nil
}
