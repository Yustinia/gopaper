package gopaper

import "encoding/json"

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

	// individual
	WallUploader Uploader `json:"uploader"`
	WallTags     []Tags   `json:"tags"`
}

type Uploader struct {
	Username string `json:"username"`
	Group    string `json:"group"`
}

type Tags struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	Date       string `json:"created_at"`
}

type Meta struct {
	CurrentPage int             `json:"current_page"`
	LastPage    int             `json:"last_page"`
	PerPage     string          `json:"per_page"`
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

type WallpaperResponse struct {
	Wall Wallpaper `json:"data"`
}

type SearchParams struct {
	KeySearch  string
	Categories string
	Purity     string
	Sorting    string
	Order      string
	AtLeast    string
	Resolution string
	Ratios     string
	Page       int
	Seed       string
}
