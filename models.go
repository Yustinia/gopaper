package gopaper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Thumbs holds the thumbnail paths of the wallpaper
type Thumbs struct {
	Large string `json:"large"`
	Orig  string `json:"original"`
	Small string `json:"small"`
}

// Wallpaper holds the data of a wallpaper
type Wallpaper struct {
	Thumbnails Thumbs `json:"thumbs"`
	ID         string `json:"id"`
	URL        string `json:"url"`
	ShortURL   string `json:"short_url"`
	Views      int    `json:"views"`
	Favorites  int    `json:"favorites"`
	// Purity indicates content rating: "sfw", "sketchy", "nsfw"
	Purity string `json:"purity"`
	// Category indiacates content: "general", "anime", "people"
	Category   string `json:"category"`
	AxisX      int    `json:"dimension_x"`
	AxisY      int    `json:"dimension_y"`
	Resolution string `json:"resolution"`
	Date       string `json:"created_at"`
	Ratio      string `json:"ratio"`
	FileSize   int    `json:"file_size"`
	FileType   string `json:"file_type"`
	// Path holds the direct image path
	Path string `json:"path"`

	WallpaperUploader Uploader `json:"uploader"`
	WallpaperTags     []Tags   `json:"tags"`
}

// Uploader holds the user to uploads the wallpaper
type Uploader struct {
	Username string `json:"username"`
	Group    string `json:"group"`
}

// Tags holds the tags assigned to the wallpaper
type Tags struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	CategoryID int    `json:"category_id"`
	Category   string `json:"category"`
	Purity     string `json:"purity"`
	Date       string `json:"created_at"`
}

// Meta holds pagindation data
type Meta struct {
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	PerPage     int    `json:"per_page"`
	Total       int    `json:"total"`
	Seed        string `json:"seed"`
}

// SearchResponse holds Meta and Wallpaper slices
type SearchResponse struct {
	Metadata   Meta        `json:"meta"`
	Wallpapers []Wallpaper `json:"data"`
}

// WallpaperResponse holds data of a singular wallpaper
type WallpaperResponse struct {
	Wall Wallpaper `json:"data"`
}

// TagResponse holds data when looking up tag IDs
type TagResponse struct {
	Tagdata Tags `json:"data"`
}

// SearchParams holds the seacrh configuration
type SearchParams struct {
	// KeySearch holds the query of what will be searched
	KeySearch string

	// Categories filters the search to "general", "anime", and "people" as indicated by: "100", "010", "001"
	Categories string

	// Purity filters the search to "sfw", "sketchy", and "nsfw" as indicated by: "100", "010", "001"
	Purity string

	// Sorting: "date_added", "relevance", "random", "views", "favorites", "toplist"
	Sorting string

	// Order sorts by "desc" or "asc"
	Order string

	// AtLeast defines the minimum resolution allowed for the search: "1920x1080"
	AtLeast string

	// Resolution defines a list of allowed resolutions: "1920x1080,2400x1080"
	Resolution string

	// Ratios defines a list of allowed ratios: "16x9,16x10"
	Ratios string

	// Page specifies which page to obtain wallpapers
	Page int

	// Seed specifies the randomness if Sorting is set to "random"
	Seed string
}

func (m *Meta) UnmarshalJSON(data []byte) error {
	type Alias struct {
		CurrentPage int             `json:"current_page"`
		LastPage    int             `json:"last_page"`
		PerPage     json.RawMessage `json:"per_page"`
		Total       int             `json:"total"`
		Seed        string          `json:"seed"`
	}

	var a Alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	m.CurrentPage = a.CurrentPage
	m.LastPage = a.LastPage
	m.Total = a.Total
	m.Seed = a.Seed

	var perPageInt int
	if err := json.Unmarshal(a.PerPage, &perPageInt); err == nil {
		m.PerPage = perPageInt
		return nil
	}

	var perPageStr string
	if err := json.Unmarshal(a.PerPage, &perPageStr); err != nil {
		return fmt.Errorf("per_page: unexpected type: %w", err)
	}

	n, err := strconv.Atoi(perPageStr)
	if err != nil {
		return fmt.Errorf("per_page: cannot convert %q to int: %w", perPageStr, err)
	}

	m.PerPage = n
	return nil
}
