package regions

// Region is a struct that represents a region available in Raider.IO API
type Region struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	ShortName string `json:"short_name"`
}

// List of regions which can be used in queries in the library
var (
	WORLD = &Region{
		Name:      "World",
		Slug:      "world",
		ShortName: "world",
	}
	US = &Region{
		Name:      "US",
		Slug:      "us",
		ShortName: "us",
	}
	EU = &Region{
		Name:      "EU",
		Slug:      "eu",
		ShortName: "eu",
	}
	KR = &Region{
		Name:      "KR",
		Slug:      "kr",
		ShortName: "kr",
	}
	TW = &Region{
		Name:      "TW",
		Slug:      "tw",
		ShortName: "tw",
	}
	CN = &Region{
		Name:      "CN",
		Slug:      "cn",
		ShortName: "cn",
	}
)
