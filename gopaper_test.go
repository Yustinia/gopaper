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
	client := NewClient()
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
