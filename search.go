package gopaper

// NewSearch returns SearchParams with sensible and safe defaults:
// Defaults to SFW content rating
func NewSearch() SearchParams {
	return SearchParams{
		Purity:     "100",
		Categories: "010",
		Sorting:    "date_added",
		Page:       1,
	}
}
