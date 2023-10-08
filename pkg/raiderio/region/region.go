package region

// Region is a struct that represents a region available in Raider.IO API
type Region struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"short_name"`
}

// Constants for regions available in Raider.IO API
const (
	WORLD string = "world"
	US    string = "us"
	EU    string = "eu"
	KR    string = "kr"
	TW    string = "tw"
	CN    string = "cn"
)
