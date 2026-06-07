package gopaper

import (
	"encoding/json"
	"os"
	"testing"
)

var APIKEY string = os.Getenv("WALLHAVEN_API_KEY")

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client.BaseURL == "" {
		t.Error("expected baseURL set, got empty string")
	}

	logPretty(t, client)
}

func TestNewClientKey(t *testing.T) {
	client := NewClientWithKey(APIKEY)

	if client.APIKey != APIKEY {
		t.Errorf("different key got, client: %s key: %s", client.APIKey, APIKEY)
	}

	logPretty(t, client)
}

func TestSFWSearch(t *testing.T) {
	client := NewClient()
	params := NewSearch()
	params.Purity = "100"
	params.Categories = "111"
	params.Page = 1

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(result.Wallpapers) == 0 {
		t.Errorf("expected wallpapers, got none")
	}
}

func TestNSFWSearch(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.Purity = "001"
	params.Categories = "010"
	params.Page = 1

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(result.Wallpapers) == 0 {
		t.Errorf("expected wallpapers, got none")
	}
}

func TestTopRange(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.TopRange = "1y"
	params.Sorting = "toplist"
	params.Order = "desc"
	params.Purity = "111"
	params.Categories = "010"

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	for _, wall := range result.Wallpapers {
		t.Log(wall.Path)
	}
}
func TestGetWallpaperDetails(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.Purity = "001"
	params.Categories = "010"
	params.Page = 1

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	details, err := client.GetWallpaperDetails(result.Wallpapers[0].ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if details.Wall.ID != result.Wallpapers[0].ID {
		t.Errorf("expected %s, got %s", details.Wall.ID, result.Wallpapers[0].ID)
	}

	logPretty(t, details.Wall)
}

func TestNextPage(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.Purity = "001"
	params.Categories = "010"
	params.Page = 5

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected the same page, got different page")
	}

	logPretty(t, result.Metadata)

	result, err = client.NextPage(result, &params)
	if err != nil {
		t.Fatalf("expected to page next, got: %v", err)
	}
	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected %v, got %v", params.Page, result.Metadata.CurrentPage)
	}

	logPretty(t, result.Metadata)
}

func TestPrevPage(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.Purity = "001"
	params.Categories = "010"
	params.Page = 5

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected the same page, got different page")
	}

	logPretty(t, result.Metadata)

	result, err = client.PrevPage(result, &params)
	if err != nil {
		t.Fatalf("expected to page previous, got: %v", err)
	}
	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected %v, got %v", params.Page, result.Metadata.CurrentPage)
	}

	logPretty(t, result.Metadata)
}
func TestSetPage(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.Purity = "001"
	params.Categories = "010"
	params.Page = 5

	result, err := client.Search(params)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected the same page, got different page")
	}

	logPretty(t, result.Metadata)

	result, err = client.SetPage(result, &params, 12)
	if err != nil {
		t.Fatalf("expected to fetch page, got: %v", err)
	}
	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected to page %v, got %v", params.Page, result.Metadata.CurrentPage)
	}

	logPretty(t, result.Metadata)
}

func TestTagLookup(t *testing.T) {
	client := NewClient()

	result, err := client.GetTagDetails(344)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	logPretty(t, result.Tagdata)
}

func TestUserSettings(t *testing.T) {
	client := NewClientWithKey(APIKEY)

	result, err := client.GetSettings()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	logPretty(t, result.UserSettings)
}

func TestCollections(t *testing.T) {
	client := NewClientWithKey(APIKEY)

	result, err := client.GetCollections()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	logPretty(t, result.UserCollections)
}

func TestPageFetching(t *testing.T) {
	client := NewClientWithKey(APIKEY)
	params := NewSearch()
	params.Purity = "111"
	params.Categories = "010"
	params.Sorting = "random"
	params.Seed = "h232d"

	result, err := client.FetchPages(&params, 2, 4)
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	for i, wall := range result {
		t.Logf("[%d] %s", i, wall.Path)
	}

	t.Logf("Wallpaper count is: %d", len(result))
}

func logPretty(t *testing.T, v any) {
	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("failed to pretty indent: %v", err)
	}

	t.Logf("\n%s", pretty)
}
