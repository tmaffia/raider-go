# Raider.io API Go wrapper

[![Go Reference](https://pkg.go.dev/badge/github.com/tmaffia/raiderio.svg)](https://pkg.go.dev/github.com/tmaffia/raiderio)
![Go Build & Test](https://github.com/tmaffia/raiderio/actions/workflows/go.yml/badge.svg)


Wrapper for the raider.io API written in Go 

## Usage

```go
client, err := raiderio.NewClient()
cq := raiderio.CharacterQuery{
	Region:        "us",
	Realm:         "illidan",
	Name:          "highervalue",
	TalentLoadout: true,
}

profile, err := client.GetCharacter(&cq)
```
