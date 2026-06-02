package gopaper

func NewSearch() SearchParams {
	return SearchParams{
		Purity:  "100",
		Sorting: "date_added",
		Page:    1,
	}
}
