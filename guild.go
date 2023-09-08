package raiderio

import "errors"

type GuildQuery struct {
	region string
	realm  string
	name   string
	fields []string
}

type GuildProfile struct {
	Name       string `json:"name"`
	Faction    string `json:"faction"`
	Region     string `json:"region"`
	Realm      string `json:"realm"`
	ProfileUrl string `json:"profile_url"`
}

func NewGuildQuery(
	region string,
	realm string,
	name string,
	fields *[]string) (*GuildQuery, error) {

	if region == "" {
		return nil, errors.New("region error")
	}

	if realm == "" {
		return nil, errors.New("realm error")
	}

	if name == "" {
		return nil, errors.New("name error")
	}

	gq := GuildQuery{
		region: region,
		realm:  realm,
		name:   name,
	}

	if fields != nil {
		gq.fields = *fields
	}

	return &gq, nil
}
