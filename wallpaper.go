package gopaper

// Return thumbnail sizes
func (w Wallpaper) ThumbSmall() string {
	return w.Thumbnails.Small
}

func (w Wallpaper) ThumbLarge() string {
	return w.Thumbnails.Large
}

func (w Wallpaper) ThumbOrig() string {
	return w.Thumbnails.Orig
}

// Return wallhaven URLs
func (w Wallpaper) WallURL() string {
	return w.URL
}

func (w Wallpaper) WallShortURL() string {
	return w.ShortURL
}

// Return direct image URL
func (w Wallpaper) ImageURL() string {
	return w.Path
}

// Purity checks
func (w Wallpaper) IsSFW() bool {
	return w.Purity == "sfw"
}

func (w Wallpaper) IsSketchy() bool {
	return w.Purity == "sketchy"
}

func (w Wallpaper) IsNSFW() bool {
	return w.Purity == "nsfw"
}

// Return tag slice
func (w Wallpaper) TagNames() []string {
	tagNames := make([]string, len(w.WallTags))

	for i, tag := range w.WallTags {
		tagNames[i] = tag.Name
	}

	return tagNames
}
