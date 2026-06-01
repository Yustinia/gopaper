package internal

import "encoding/json"

type Client struct {
	APIKey  string
	BaseURL string
}

func NewClient(apiKey string) Client {
	return Client{APIKey: apiKey, BaseURL: "https://wallhaven.cc/api/v1"}
}

type Thumbs struct {
	Large string `json:"large"`
	Orig  string `json:"original"`
	Small string `json:"small"`
}

type Wallpaper struct {
	Thumbnails Thumbs `json:"thumbs"`
	ID         string `json:"id"`
	URL        string `json:"url"`
	ShortURL   string `json:"short_url"`
	Views      int    `json:"views"`
	Favorites  int    `json:"favorites"`
	Purity     string `json:"purity"`
	Category   string `json:"category"`
	Resolution string `json:"resolution"`
	Date       string `json:"created_at"`
	Ratio      string `json:"ratio"`
	Path       string `json:"path"`
}

type Meta struct {
	CurrentPage int             `json:"current_page"`
	LastPage    int             `json:"last_page"`
	PerPage     int             `json:"per_page"`
	Total       int             `json:"total"`
	Seed        string          `json:"seed"`
	SearchQuery json.RawMessage `json:"query"`
}

type Query struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

type SearchResponse struct {
	Metadata Meta        `json:"meta"`
	Wall     []Wallpaper `json:"data"`
}
