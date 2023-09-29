package raiderio

import "errors"

// GuildQuery is a struct that represents the query parameters
// sent for a guild profile request
type GuildQuery struct {
	region string
	realm  string
	name   string
	fields []string
}

// GuildProfile is a struct that represents the response from
// a guild profile request
type GuildProfile struct {
	Name       string `json:"name"`
	Faction    string `json:"faction"`
	Region     string `json:"region"`
	Realm      string `json:"realm"`
	ProfileUrl string `json:"profile_url"`
}

// NewGuildQuery creates a new GuildQuery struct
// It returns an error if any of the required parameters are empty
// or if the fields parameter is not nil and is empty
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
