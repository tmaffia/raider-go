package raiderio

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/tmaffia/raiderio/pkg/raiderio/expansion"
)

// Base URL for the Raider.IO API
const baseUrl string = "https://raider.io/api"

// Client is the main struct for interacting with the Raider.IO API
type Client struct {
	ApiUrl     string
	HttpClient *http.Client
}

// NewClient creates a new Client struct
func NewClient() *Client {
	var c Client
	c.ApiUrl = baseUrl + "/v1"
	c.HttpClient = &http.Client{}
	return &c
}

// GetCharacter retrieves a character profile from the Raider.IO API
// It returns an error if the API returns a non-200 status code, or if the
// response body cannot be read or mapped to the CharacterProfile struct
func (c *Client) GetCharacter(cq *CharacterQuery) (*Character, error) {
	err := validateCharacterQuery(cq)
	if err != nil {
		return nil, err
	}

	reqUrl := c.ApiUrl + "/characters/profile?region=" + cq.Region.Slug + "&realm=" + cq.Realm + "&name=" + cq.Name
	if cq.fields != nil && len(cq.fields) != 0 {
		reqUrl += "&fields=" + strings.Join(cq.fields, ",")
	}

	body, err := c.getAPIResponse(reqUrl)
	if err != nil {
		return nil, err
	}

	var profile Character
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return nil, errors.New("error unmarshalling character profile")
	}

	return &profile, nil
}

// GetGuild retrieves a guild profile from the Raider.IO API
// It returns an error if the API returns a non-200 status code, or if the
// response body cannot be read or mapped to the GuildProfile struct
func (c *Client) GetGuild(gq *GuildQuery) (*Guild, error) {
	err := createGuildQuery(gq)
	if err != nil {
		return nil, err
	}

	reqUrl := c.ApiUrl + "/guilds/profile?region=" + gq.Region.Slug + "&realm=" + gq.Realm + "&name=" + gq.Name
	if gq.fields != nil && len(gq.fields) != 0 {
		reqUrl += "&fields=" + strings.Join(gq.fields, ",")
	}

	body, err := c.getAPIResponse(reqUrl)
	if err != nil {
		return nil, err
	}

	profile, err := unmarshalGuild(body)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// GetRaids retrieves a list of raids from the Raider.IO API
// It returns an error if the API returns a non-200 status code, or if the
// response body cannot be read or mapped to the Raids struct
// Takes an Expansion enum as a parameter
func (c *Client) GetRaids(e expansion.Expansion) (*Raids, error) {
	reqUrl := c.ApiUrl + "/raiding/static-data?expansion_id=" + fmt.Sprintf("%d", e)
	body, err := c.getAPIResponse(reqUrl)
	if err != nil {
		return nil, err
	}

	var raids Raids
	err = json.Unmarshal(body, &raids)
	if err != nil {
		return nil, errors.New("error unmarshalling raids")
	}

	return &raids, nil
}

// GetRaidRankings retrieves a list of raid rankings from the Raider.IO API
// It returns an error if the API returns a non-200 status code, or if the
// response body cannot be read or mapped to the RaidRankings struct
// Takes a RaidQuery struct as a parameter
func (c *Client) GetRaidRankings(rq *RaidQuery) (*RaidRankings, error) {
	err := validateRaidRankingsQuery(rq)
	if err != nil {
		return nil, err
	}

	reqUrl := c.ApiUrl + "/raiding/raid-rankings?raid=" + rq.Slug +
		"&difficulty=" + string(rq.Difficulty) + "&region=" + rq.Region.Slug

	if rq.Realm != "" {
		reqUrl += "&realm=" + rq.Realm
	}

	if rq.Limit != 0 {
		reqUrl += "&limit=" + fmt.Sprintf("%d", rq.Limit)
	}

	if rq.Page != 0 {
		reqUrl += "&page=" + fmt.Sprintf("%d", rq.Page)
	}

	body, err := c.getAPIResponse(reqUrl)
	if err != nil {
		return nil, err
	}

	var rankings RaidRankings
	err = json.Unmarshal(body, &rankings)
	if err != nil {
		return nil, errors.New("error unmarshalling raid rankings")
	}

	return &rankings, nil
}
