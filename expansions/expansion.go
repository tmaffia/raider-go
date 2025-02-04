package expansions

// Expansion Type is an enum that represents an expansion
type Expansion int

// Constants for expansions available in Raider.IO API
// Expansions are referenced as ints in the API
const (
	WarWithin        Expansion = 10
	Dragonflight     Expansion = 9
	Shadowlands      Expansion = 8
	BattleForAzeroth Expansion = 7
	Legion           Expansion = 6
)
