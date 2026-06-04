package gopaper

import (
	"os"
	"testing"
)

var APIKEY string = os.Getenv("WALLHAVEN_API_KEY")

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client.BaseURL == "" {
		t.Error("expected baseURL set, got empty string")
	}
}

func TestNewClientKey(t *testing.T) {
	client := NewClientWithKey(APIKEY)

	if client.APIKey != APIKEY {
		t.Errorf("different key got, client: %s key: %s", client.APIKey, APIKEY)
	}
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

	result, err = client.NextPage(result, &params)
	if err != nil {
		t.Fatalf("expected to page next, got: %v", err)
	}
	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected %v, got %v", params.Page, result.Metadata.CurrentPage)
	}
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

	result, err = client.PrevPage(result, &params)
	if err != nil {
		t.Fatalf("expected to page previous, got: %v", err)
	}
	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected %v, got %v", params.Page, result.Metadata.CurrentPage)
	}
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

	result, err = client.SetPage(result, &params, 12)
	if err != nil {
		t.Fatalf("expected to fetch page, got: %v", err)
	}
	if result.Metadata.CurrentPage != params.Page {
		t.Errorf("expected to page %v, got %v", params.Page, result.Metadata.CurrentPage)
	}
}

func TestTagLookup(t *testing.T) {
	client := NewClient()

	result, err := client.GetTagDetails(344)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	t.Logf("ID: %v", result.Tagdata.TagID())
	t.Logf("Name: %v", result.Tagdata.TagName())
	t.Logf("Alias: %v", result.Tagdata.TagAlias())
	t.Logf("CategoryID: %v", result.Tagdata.TagCategoryID())
	t.Logf("Category%v", result.Tagdata.TagCategory())
	t.Logf("Purity: %v", result.Tagdata.TagPurity())
	t.Logf("Date: %v", result.Tagdata.TagDate())
}

func TestUserSettings(t *testing.T) {
	client := NewClientWithKey(APIKEY)

	result, err := client.GetSettings()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	t.Logf("%+v", result)
}

func TestCollections(t *testing.T) {
	client := NewClientWithKey(APIKEY)

	result, err := client.GetCollections()
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	t.Logf("%+v", result)
}
