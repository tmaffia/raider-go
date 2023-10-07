package util

// Expansion Type is an enum that represents an expansion
type Expansion int

const (
	Dragonflight     Expansion = 9
	Shadowlands      Expansion = 8
	BattleForAzeroth Expansion = 7
	Legion           Expansion = 6
)

// Region is a struct that represents a region available in Raider.IO API
type Region struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"short_name"`
}

// Realm is a struct that represents a realm available in Raider.IO API
type Realm struct {
	Id               int    `json:"id"`
	ConnectedRealmId int    `json:"connectedRealmId"`
	Name             string `json:"name"`
	AltName          string `json:"altName"`
	Slug             string `json:"slug"`
	AltSlug          string `json:"altSlug"`
	Locale           string `json:"locale"`
	IsConnected      bool   `json:"isConnected"`
}
