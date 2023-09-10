package raiderio

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const baseUrl string = "https://raider.io/api"

type Client struct {
	apiUrl     string
	httpClient *http.Client
}

func NewClient() (*Client, error) {
	var c Client
	c.apiUrl = baseUrl + "/v1"
	c.httpClient = &http.Client{}
	return &c, nil
}

func (c *Client) GetCharacterProfile(cq *CharacterQuery) (*CharacterProfile, error) {
	reqUrl := c.apiUrl + "/characters/profile?region=" + cq.Region + "&realm=" + cq.Realm + "&name=" + cq.Name
	if cq.Fields != nil && len(cq.Fields) != 0 {
		reqUrl += "&fields=" + strings.Join(cq.Fields, ",")
	}

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

	var profile CharacterProfile
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return nil, errors.New("character profile mapping error")
	}

	return &profile, nil
}
