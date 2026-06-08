package gopaper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Thumbs contain thumbnail URLs of a wallpaper
type Thumbs struct {
	Large string `json:"large"`
	Orig  string `json:"original"`
	Small string `json:"small"`
}

// Wallpaper represents a single wallpaper from the API
type Wallpaper struct {
	Thumbnails Thumbs `json:"thumbs"`
	// ID contains the uniquely assigned ID for the wallpaper
	ID string `json:"id"`
	// URL contains the full wallhaven URL
	URL string `json:"url"`
	// ShortURL contains the shortened wallhaven URL
	ShortURL string `json:"short_url"`
	// Views is the number of times a wallpaper has ben seen
	Views int `json:"views"`
	// Favorites is the number of users that has favorited a wallpaper
	Favorites int `json:"favorites"`
	// Purity indicates rating: "sfw", "sketchy", "nsfw"
	Purity string `json:"purity"`
	// Category indicates content: "general", "anime", "people"
	Category string `json:"category"`
	// AxisX contains the width pixel count
	AxisX int `json:"dimension_x"`
	// AxisY contains the height pixel count
	AxisY int `json:"dimension_y"`
	// Resolution contains the WIDTH x HEIGHT pixel count
	Resolution string `json:"resolution"`
	// Date indicates when the wallpaper is was uploaded
	Date string `json:"created_at"`
	// Ratio is the wallpaper's aspect ratio
	Ratio string `json:"ratio"`
	// FileSize is how large the file is in Bytes
	FileSize int `json:"file_size"`
	// FileType is the format the wallpaper
	FileType string `json:"file_type"`
	// Path holds the direct image path
	Path string `json:"path"`

	WallpaperUploader Uploader `json:"uploader"`
	WallpaperTags     []Tags   `json:"tags"`
}

// Uploader contains information who uploaded the wallpaper
type Uploader struct {
	// Username is the name of the uploader
	Username string `json:"username"`
	// Group is where the uploader is a part of
	Group string `json:"group"`
}

// Tags represent a tag assigned to a wallpaper.
type Tags struct {
	// ID is the unique tag identifier.
	ID int `json:"id"`
	// Name is the display name of the tag.
	Name string `json:"name"`
	// Alias is an alternative name or shorthand for the tag, if any.
	Alias string `json:"alias"`
	// CategoryID is the identifier of the tag's parent category.
	CategoryID int `json:"category_id"`
	// Category is the human-readable name of the parent category.
	Category string `json:"category"`
	// Purity is the content rating associated with this tag: "sfw", "sketchy", or "nsfw".
	Purity string `json:"purity"`
	// Date is the RFC3339 timestamp when the tag was created.
	Date string `json:"created_at"`
}

// Meta contains pagination metadata for search results.
type Meta struct {
	// CurrentPage is the page number returned in this response.
	CurrentPage int `json:"current_page"`
	// LastPage is the final available page for this query.
	LastPage int `json:"last_page"`
	// PerPage is the number of results returned per page.
	PerPage int `json:"per_page"`
	// Total is the total number of results matching the query.
	Total int `json:"total"`
	// Seed is the randomization seed used for "random" sorting.
	// Reuse this seed to retrieve the same result set again.
	Seed string `json:"seed"`
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

// SettingsResponse holds user settings data
type SettingsResponse struct {
	UserSettings Settings `json:"data"`
}

// Settings represents a user's Wallhaven account preferences.
type Settings struct {
	// ThumbSize is the preferred thumbnail resolution (e.g., "small", "large").
	ThumbSize string `json:"thumb_size"`
	// PerPage is the number of wallpapers to display per page.
	PerPage string `json:"per_page"`
	// Purity is the list of enabled content ratings (e.g., ["sfw", "sketchy"]).
	Purity []string `json:"purity"`
	// Categories is the list of enabled content categories (e.g., ["general", "anime"]).
	Categories []string `json:"categories"`
	// Resolutions is the list of allowed wallpaper resolutions.
	Resolutions []string `json:"resolutions"`
	// Ratios is the list of preferred aspect ratios.
	Ratios []string `json:"aspect_ratios"`
	// ToplistRange is the time period for toplist results (e.g., "1d", "1w", "1M", "1y").
	ToplistRange string `json:"toplist_range"`
	// TagBlacklist is the list of tags to exclude from search results.
	TagBlacklist []string `json:"tag_blacklist"`
	// UserBlacklist is the list of uploaders whose wallpapers are hidden.
	UserBlacklist []string `json:"user_blacklist"`
}

// CollectionResponse holds wallpaper collections
type CollectionResponse struct {
	UserCollections []Collections `json:"data"`
}

// Collections represents a user's wallpaper collection.
type Collections struct {
	// ID is the unique collection identifier.
	ID int `json:"id"`
	// Label is the collection's display name.
	Label string `json:"label"`
	// Views is the number of times the collection has been viewed.
	Views int `json:"views"`
	// Public indicates whether the collection is publicly visible (1) or private (0).
	Public int `json:"public"`
	// Count is the number of wallpapers in the collection.
	Count int `json:"count"`
}

// SearchParams configures a wallpaper search query.
type SearchParams struct {
	// KeySearch is the search keyword or phrase.
	KeySearch string

	// Categories filters by content type using a 3-bit string:
	// "100" = general, "010" = anime, "001" = people.
	// Combine with "1" for enabled (e.g., "111" for all).
	Categories string

	// Purity filters by content rating using a 3-bit string:
	// "100" = sfw, "010" = sketchy, "001" = nsfw.
	// Combine with "1" for enabled (e.g., "110" for sfw + sketchy).
	Purity string

	// Sorting determines result order: "date_added", "relevance", "random",
	// "views", "favorites", or "toplist".
	Sorting string

	// Order is the sort direction: "desc" or "asc".
	Order string

	// TopRange filters results to a time window when Sorting is "toplist":
	// "1d", "3d", "1w", "1M", "3M", "6M", "1y".
	TopRange string

	// AtLeast is the minimum resolution (e.g., "1920x1080").
	AtLeast string

	// Resolution is a comma-separated list of exact resolutions
	// (e.g., "1920x1080,2400x1080").
	Resolution string

	// Ratios is a comma-separated list of aspect ratios
	// (e.g., "16x9,16x10").
	Ratios string

	// Page is the result page to retrieve (1-based).
	Page int

	// Seed is the randomization seed for consistent results when
	// Sorting is "random". Reuse the same seed to get the same page.
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
