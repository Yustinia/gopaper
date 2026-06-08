# GoPaper

[![GO documentation](https://pkg.go.dev/badge/github.com/Yustinia/gopaper.svg)](https://pkg.go.dev/github.com/Yustinia/gopaper)

A Go library for interacting with the [Wallhaven](https://wallhaven.cc/) API to search and retrieve wallpapers.

## Installation

Installing the wrapper for use is as simple as:

```bash
go get github.com/Yustinia/gopaper
```

## Usage

Quickstart:

```go
package main

import (
    "fmt"
    "github.com/Yustinia/gopaper"
)

func main() {
    // Create a client without API Key
    client := gopaper.NewClient()

    // Use this if you have a valid API key
    // client := gopaper.NewClientWithKey("APIKey")

    // Create new search
    params := gopaper.NewSearch()
    // Filter content to Anime
    params.Categories = "010"
    // Filter content rating to SFW
    params.Purity = "100"

    // Configure the search parameters
    params.KeySearch = "japan"

    // Perform the search and provide
    result, err := client.Search(params)
    if err != nil {
        panic(err)
    }

    for i, wall := range result.Wallpapers {
        fmt.Printf("[%d] %s\n", i, wall.Path)
    }

    // Get full details for a specific wallpaper from the result
    details, err := client.GetWallpaperDetails(result.Wallpapers[0].ID)
    // or by providing the ID as the argument
    // details, err := client.GetWallpaperDetails("poyzl3")
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
