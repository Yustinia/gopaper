package gopaper

// Return the path of the small thumbnail
func (w Wallpaper) ThumbSmall() string {
	return w.Thumbnails.Small
}

// Returns the path of the large thumbnail
func (w Wallpaper) ThumbLarge() string {
	return w.Thumbnails.Large
}

// Returns the original size of the thumbnail
func (w Wallpaper) ThumbOrig() string {
	return w.Thumbnails.Orig
}

// Returns the full wallhaven link
func (w Wallpaper) WallURL() string {
	return w.URL
}

// Returns the shortened wallhaven link
func (w Wallpaper) WallShortURL() string {
	return w.ShortURL
}

// Returns a direct image URL
func (w Wallpaper) ImageURL() string {
	return w.Path
}

// IsSFW checks if wallpaper content rating is SFW
func (w Wallpaper) IsSFW() bool {
	return w.Purity == "sfw"
}

// IsSketchy checks if wallpaper content rating is Sketchy
func (w Wallpaper) IsSketchy() bool {
	return w.Purity == "sketchy"
}

// IsNSFW checks if wallpaper content rating is NSFW
func (w Wallpaper) IsNSFW() bool {
	return w.Purity == "nsfw"
}

// IsGeneral checks if wallpaper category rating is General
func (w Wallpaper) IsGeneral() bool {
	return w.Category == "general"
}

// IsAnime checks if wallpaper category is Anime
func (w Wallpaper) IsAnime() bool {
	return w.Category == "anime"
}

// IsPeople checks if wallpaper category is People
func (w Wallpaper) IsPeople() bool {
	return w.Category == "people"
}

// TagNames returns a slice of tags assigned to the wallpaper
func (w Wallpaper) TagNames() []string {
	tagNames := make([]string, len(w.WallpaperTags))

	for i, tag := range w.WallpaperTags {
		tagNames[i] = tag.Name
	}

	return tagNames
}

// Returns pixel count at the X axis
func (w Wallpaper) WallXAxis() int {
	return w.AxisX
}

// Returns pixel count at the Y axis
func (w Wallpaper) WallYAxis() int {
	return w.AxisY
}

// Return file size in bytes
func (w Wallpaper) SizeByte() int {
	return w.FileSize
}

// Return file size in MiB
func (w Wallpaper) SizeMiB() float64 {
	return float64(w.FileSize) / 1024 / 1024
}

// Return file size in KiB
func (w Wallpaper) SizeKiB() float64 {
	return float64(w.FileSize) / 1024
}

// Return file type in "image/type"
func (w Wallpaper) WallFileType() string {
	return w.FileType
}
