package gopaper

type SettingsResponse struct {
	UserSettings Settings `json:"data"`
}

type Settings struct {
	ThumbSize     string   `json:"thumb_size"`
	PerPage       string   `json:"per_page"`
	Purity        []string `json:"purity"`
	Categories    []string `json:"categories"`
	Resolutions   []string `json:"resolutions"`
	Ratios        []string `json:"aspect_ratios"`
	ToplistRange  string   `json:"toplist_range"`
	TagBlacklist  []string `json:"tag_blacklist"`
	UserBlacklist []string `json:"user_blacklist"`
}

type CollectionResponse struct {
	UserCollections []Collections `json:"data"`
}

type Collections struct {
	ID     int    `json:"id"`
	Label  string `json:"label"`
	Views  int    `json:"views"`
	Public int    `json:"public"`
	Count  int    `json:"count"`
}
