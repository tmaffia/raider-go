package raiderio

import "errors"

type CharacterQuery struct {
	Region string   `json:"region"`
	Realm  string   `json:"realm"`
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
}

type CharacterProfile struct {
	Name              string        `json:"name"`
	Race              string        `json:"race"`
	Class             string        `json:"class"`
	ActiveSpec        string        `json:"active_spec_name"`
	ActiveRole        string        `json:"active_spec_role"`
	Gender            string        `json:"gender"`
	Faction           string        `json:"faction"`
	AchievementPoints int64         `json:"achievement_points"`
	HonorableKills    int64         `json:"honorable_kills"`
	ThumbnailUrl      string        `json:"thumbnail_url"`
	Region            string        `json:"region"`
	Realm             string        `json:"realm"`
	LastCrawledAt     string        `json:"last_crawled_at"`
	ProfileUrl        string        `json:"profile_url"`
	ProfileBanner     string        `json:"profile_banner"`
	TalentLoadout     TalentLoadout `json:"talentLoadout"`
	Gear              Gear          `json:"gear"`
}

type Gear struct {
	UpdatedAt         string `json:"updated_at"`
	ItemLevelEquipped int    `json:"item_level_equipped"`
	ItemLevelTotal    int    `json:"item_level_total"`
	Items             Items  `json:"items"`
}

type Items struct {
	Head     Item `json:"head"`
	Neck     Item `json:"neck"`
	Shoulder Item `json:"shoulder"`
	Back     Item `json:"back"`
	Chest    Item `json:"chest"`
	Wrist    Item `json:"wrist"`
	Hands    Item `json:"hands"`
	Waist    Item `json:"waist"`
	Legs     Item `json:"legs"`
	Feet     Item `json:"feet"`
	Finger1  Item `json:"finger1"`
	Finger2  Item `json:"finger2"`
	Trinket1 Item `json:"trinket1"`
	Trinket2 Item `json:"trinket2"`
	Mainhand Item `json:"mainhand"`
	Offhand  Item `json:"offhand"`
	Shirt    Item `json:"shirt"`
	Tabard   Item `json:"tabard"`
}

type Item struct {
	ID          int    `json:"item_id"`
	ItemLevel   int    `json:"item_level"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	ItemQuality int    `json:"item_quality"`
	IsLegendary bool   `json:"is_legendary"`
	Gems        []int  `json:"gems"`
	Bonuses     []int  `json:"bonuses"`
}

type TalentLoadout struct {
	LoadoutSpecID int    `json:"loadout_spec_id"`
	LoadoutText   string `json:"loadout_text"`
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
