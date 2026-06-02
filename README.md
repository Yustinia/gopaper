# GoPaper

A GO library to interact with the Wallhaven API to search and retrieve wallpapers

## Installation

```
go get github.com/Yustinia/gopaper
```

## Usage

Quickstart:

```go
func main() {
    client := gopaper.NewClient("APIKey")
    params := gopaper.NewSearch()
    params.KeySearch = "japan"

    result, err := client.Search(params)
    if err != nil {
        panic(err)
    }

    for i, wall := range result.Wal {
        fmt.Printf("[%d] %s\n", i, wall.Path)
    }
}
```

## Features

- Search Wallpapers
- Get wallpaper details
- Filter through search parameters:
  - Categories
  - Purity
  - Sorting
  - Order
  - At Least
  - Resolution
  - Ratios
  - Page
  - Seed
