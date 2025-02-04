package realms

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
