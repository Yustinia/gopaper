package gopaper

// NewSearch initializes the default search parameters
func NewSearch() SearchParams {
	return SearchParams{
		Purity:     "100",
		Categories: "010",
		Sorting:    "date_added",
		Page:       1,
	}
}
