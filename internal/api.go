package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	APIKey  string
	BaseURL string
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

func NewSearch() SearchParams {
	return SearchParams{
		Purity:  "100",
		Sorting: "date_added",
		Page:    1,
	}
}

func NewClient(apiKey string) Client {
	return Client{APIKey: apiKey, BaseURL: "https://wallhaven.cc/api/v1"}
}

func buildParams(sp SearchParams, apiKey string) url.Values {
	params := url.Values{}

	if apiKey != "" {
		params.Set("apikey", apiKey)
	}

	if sp.KeySearch != "" {
		params.Set("q", sp.KeySearch)
	}
	if sp.Categories != "" {
		params.Set("categories", sp.Categories)
	}
	if sp.Purity != "" {
		params.Set("purity", sp.Purity)
	}
	if sp.Sorting != "" {
		params.Set("sorting", sp.Sorting)
	}
	if sp.Order != "" {
		params.Set("order", sp.Order)
	}
	if sp.AtLeast != "" {
		params.Set("atleast", sp.AtLeast)
	}
	if sp.Resolution != "" {
		params.Set("resolutions", sp.Resolution)
	}
	if sp.Ratios != "" {
		params.Set("ratios", sp.Ratios)
	}
	if sp.Page != 0 {
		params.Set("page", fmt.Sprintf("%d", sp.Page))
	}
	if sp.Seed != "" {
		params.Set("seed", sp.Seed)
	}

	return params
}

func (c *Client) Search(sp SearchParams) (SearchResponse, error) {
	var result SearchResponse
	params := buildParams(sp, c.APIKey)

	buildURL := fmt.Sprintf("%s/search?%s", c.BaseURL, params.Encode())

	resp, err := http.Get(buildURL)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("something went wrong: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SearchResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("failed to decode: %w", err)
	}

	return result, nil
}

func (c *Client) GetWallpaperDetails(wallID string) (WallpaperResponse, error) {
	var result WallpaperResponse

	buildURL := fmt.Sprintf("%s/w/%s", c.BaseURL, wallID)
	if c.APIKey != "" {
		buildURL = fmt.Sprintf("%s?apikey=%s", buildURL, c.APIKey)
	}

	resp, err := http.Get(buildURL)
	if err != nil {
		return WallpaperResponse{}, fmt.Errorf("something went wrong: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WallpaperResponse{}, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return WallpaperResponse{}, fmt.Errorf("failed to decode: %w", err)
	}

	return result, nil
}
