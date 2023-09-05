package raiderio

import (
	"log"
	"testing"
)

func TestNewCharacterQuery(t *testing.T) {
	region := "us"
	realm := "illidan"
	name := "thehighvalue"
	fields := []string{"gear", "mythic_plus_rank", "mythic_plus_recent_runs"}

	cq, err := newCharacterQuery(region, realm, name, fields)

	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	log.Println(cq)
}
