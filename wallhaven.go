package gopaper

import "fmt"

func (c *Client) Search(sp SearchParams) (SearchResponse, error) {
	params := buildParams(sp, c.APIKey)

	buildURL := fmt.Sprintf("%s/search?%s", c.BaseURL, params.Encode())

	return doRequest[SearchResponse](buildURL)
}

func (c *Client) GetWallpaperDetails(wallID string) (WallpaperResponse, error) {
	params := buildBaseParams(c.APIKey)

	buildURL := fmt.Sprintf("%s/w/%s", c.BaseURL, wallID)
	if c.APIKey != "" {
		buildURL = fmt.Sprintf("%s?%s", buildURL, params.Encode())
	}

	return doRequest[WallpaperResponse](buildURL)
}
