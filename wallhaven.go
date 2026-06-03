package gopaper

import (
	"errors"
	"fmt"
)

var ErrLastPage = errors.New("already on the last page")
var ErrFirstPage = errors.New("already on first page")
var ErrInvalidPage = errors.New("not a valid page")

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
