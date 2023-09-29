# Raider.io API Go wrapper

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
