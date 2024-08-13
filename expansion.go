package raiderio

// Expansion Type is an enum that represents an expansion
type Expansion int

// Constants for expansions available in Raider.IO API
// Expansions are referenced as ints in the API
var (
	Expansions = struct {
		WarWithin        Expansion
		Dragonflight     Expansion
		Shadowlands      Expansion
		BattleForAzeroth Expansion
		Legion           Expansion
	}{
		WarWithin:        10,
		Dragonflight:     9,
		Shadowlands:      8,
		BattleForAzeroth: 7,
		Legion:           6,
	}
)
