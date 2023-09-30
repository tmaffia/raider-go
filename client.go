package raiderio

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Base URL for the Raider.IO API
const baseUrl string = "https://raider.io/api"

// Client is the main struct for interacting with the Raider.IO API
type Client struct {
	apiUrl     string
	httpClient *http.Client
}

// NewClient creates a new Client struct
func NewClient() *Client {
	var c Client
	c.apiUrl = baseUrl + "/v1"
	c.httpClient = &http.Client{}
	return &c
}

// GetCharacter retrieves a character profile from the Raider.IO API
// It returns an error if the API returns a non-200 status code, or if the
// response body cannot be read or mapped to the CharacterProfile struct
func (c *Client) GetCharacter(cq *CharacterQuery) (*Character, error) {
	err := createCharacterQuery(cq)
	if err != nil {
		return nil, err
	}

	reqUrl := c.apiUrl + "/characters/profile?region=" + cq.Region + "&realm=" + cq.Realm + "&name=" + cq.Name
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
		return nil, errors.New("character profile mapping error")
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

	reqUrl := c.apiUrl + "/guilds/profile?region=" + gq.Region + "&realm=" + gq.Realm + "&name=" + gq.Name
	if gq.fields != nil && len(gq.fields) != 0 {
		reqUrl += "&fields=" + strings.Join(gq.fields, ",")
	}

	body, err := c.getAPIResponse(reqUrl)
	if err != nil {
		return nil, err
	}
	var profile Guild
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return nil, errors.New("guild profile mapping error")
	}
	return &profile, nil
}

// getAPIResponse is a helper function that makes a GET request to the Raider.IO API
// It returns an error if the API returns a non-200 status code, or if the
// response body cannot be read
func (c *Client) getAPIResponse(reqUrl string) ([]byte, error) {
	resp, err := c.httpClient.Get(reqUrl)

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	if err != nil {
		return nil, errors.New("api response error")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("api response error")
	}
	return body, nil
}
