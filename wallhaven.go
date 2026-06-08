package gopaper

import (
	"errors"
	"fmt"
)

var ErrLastPage = errors.New("already on the last page")
var ErrFirstPage = errors.New("already on first page")
var ErrInvalidPage = errors.New("not a valid page")

var ErrAPISettings = errors.New("API required for reading settings")
var ErrAPICollections = errors.New("API required for retreiving collections")

var ErrInvalidPageRange = errors.New("invalid page range")
var ErrInvalidWallCount = errors.New("invalid wallpaper fetch count")

// Search performs the search based on provided search parameters
func (c *Client) Search(sp SearchParams) (SearchResponse, error) {
	params := buildParams(sp, c.APIKey)

	buildURL := fmt.Sprintf("%s/search?%s", c.BaseURL, params.Encode())

	return doRequest[SearchResponse](buildURL)
}

// GetWallpaperDetails retrieves metadata of a wallpaper provided with the ID
func (c *Client) GetWallpaperDetails(wallID string) (WallpaperResponse, error) {
	params := buildBaseParams(c.APIKey)

	buildURL := fmt.Sprintf("%s/w/%s?%s", c.BaseURL, wallID, params.Encode())

	return doRequest[WallpaperResponse](buildURL)
}

// NextPage retrieves new results of the next page
func (c *Client) NextPage(result SearchResponse, sp *SearchParams) (SearchResponse, error) {
	if result.Metadata.CurrentPage >= result.Metadata.LastPage {
		return SearchResponse{}, ErrLastPage
	}

	sp.Page++

	return c.Search(*sp)
}

// PrevPage retrieves results of the previous page
func (c *Client) PrevPage(result SearchResponse, sp *SearchParams) (SearchResponse, error) {
	if result.Metadata.CurrentPage <= 1 {
		return SearchResponse{}, ErrFirstPage
	}

	sp.Page--

	return c.Search(*sp)
}

// SetPage retrieves results based on the provided page
func (c *Client) SetPage(result SearchResponse, sp *SearchParams, page int) (SearchResponse, error) {
	if page > result.Metadata.LastPage {
		return SearchResponse{}, ErrLastPage
	} else if page < 1 {
		return SearchResponse{}, ErrFirstPage
	}

	sp.Page = page

	return c.Search(*sp)
}

// FetchPages returns multiple wallpaper results based on the page range set on fromPage to toPage
func (c *Client) FetchPages(sp *SearchParams, fromPage int, toPage int) ([]Wallpaper, error) {
	var fetchedWalls []Wallpaper

	if fromPage > toPage {
		return fetchedWalls, ErrInvalidPageRange
	}

	sp.Page = fromPage
	result, err := c.Search(*sp)
	if err != nil {
		return fetchedWalls, err
	}

	if toPage > result.Metadata.LastPage {
		toPage = result.Metadata.LastPage
	}

	fetchedWalls = append(fetchedWalls, result.Wallpapers...)
	sp.Page++

	for i := fromPage + 1; i <= toPage; i++ {
		sp.Page = i

		result, err := c.Search(*sp)
		if err != nil {
			return fetchedWalls, err
		}

		fetchedWalls = append(fetchedWalls, result.Wallpapers...)
	}

	return fetchedWalls, nil
}

// FetchWallpaperCount returns wallpaper items depending on the wallCount
func (c *Client) FetchWallpaperCount(sp *SearchParams, wallCount int) ([]Wallpaper, error) {
	var fetchedWalls []Wallpaper

	if wallCount <= 0 {
		return fetchedWalls, ErrInvalidWallCount
	}

	result, err := c.Search(*sp)
	if err != nil {
		return fetchedWalls, err
	}
	maxWallCount := result.Metadata.Total
	fetchedWalls = append(fetchedWalls, result.Wallpapers...)
	sp.Page++

	for {
		result, err = c.Search(*sp)
		if err != nil {
			return fetchedWalls, err
		}

		fetchedWalls = append(fetchedWalls, result.Wallpapers...)
		sp.Page++

		if len(fetchedWalls) >= wallCount || len(fetchedWalls) >= maxWallCount {
			break
		}
	}
	if len(fetchedWalls) > wallCount {
		fetchedWalls = fetchedWalls[:wallCount]
	}

	return fetchedWalls, nil
}

// GetTagDetails retrieves details of a tag based on the tag ID
func (c *Client) GetTagDetails(tagID int) (TagResponse, error) {
	params := buildBaseParams(c.APIKey)

	buildURL := fmt.Sprintf("%s/tag/%d?%s", c.BaseURL, tagID, params.Encode())

	return doRequest[TagResponse](buildURL)
}

// GetSettings retrieves user settings data
func (c *Client) GetSettings() (SettingsResponse, error) {
	if c.APIKey == "" {
		return SettingsResponse{}, ErrAPISettings
	}

	params := buildBaseParams(c.APIKey)

	buildURL := fmt.Sprintf("%s/settings?%s", c.BaseURL, params.Encode())

	return doRequest[SettingsResponse](buildURL)
}

// GetCollections retrieves user collection data
func (c *Client) GetCollections() (CollectionResponse, error) {
	if c.APIKey == "" {
		return CollectionResponse{}, ErrAPICollections
	}

	params := buildBaseParams(c.APIKey)

	buildURL := fmt.Sprintf("%s/collections?%s", c.BaseURL, params.Encode())

	return doRequest[CollectionResponse](buildURL)
}
