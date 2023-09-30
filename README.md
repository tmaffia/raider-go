# Raider.io API Go wrapper

[![Go Reference](https://pkg.go.dev/badge/github.com/tmaffia/raiderio.svg)](https://pkg.go.dev/github.com/tmaffia/raiderio)
![Go Build & Test](https://github.com/tmaffia/raiderio/actions/workflows/go.yml/badge.svg)


Raider go is a wrapper for the raider.io API written in Go 

## Usage

```go
client, err := raiderio.NewClient()
query, err := raiderio.NewCharacterQuery(
    "us",
    "illidan",
    "Highervalue",
    nil,
)

profile, err := client.GetCharacterProfile(query)
```
