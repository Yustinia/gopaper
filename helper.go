package gopaper

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func buildBaseParams(apiKey string) url.Values {
	params := url.Values{}
	if apiKey != "" {
		params.Set("apikey", apiKey)
	}

	return params
}

func setIfNotEmpty(v url.Values, key string, value string) {
	if value != "" {
		v.Set(key, value)
	}
}

func buildParams(sp SearchParams, apiKey string) url.Values {
	params := buildBaseParams(apiKey)

	setIfNotEmpty(params, "q", sp.KeySearch)
	setIfNotEmpty(params, "categories", sp.Categories)
	setIfNotEmpty(params, "purity", sp.Purity)
	setIfNotEmpty(params, "sorting", sp.Sorting)
	setIfNotEmpty(params, "order", sp.Order)
	setIfNotEmpty(params, "atleast", sp.AtLeast)
	setIfNotEmpty(params, "resolutions", sp.Resolution)
	setIfNotEmpty(params, "ratios", sp.Ratios)
	setIfNotEmpty(params, "seed", sp.Seed)

	if sp.Page != 0 {
		params.Set("page", strconv.Itoa(sp.Page))
	}

	return params
}

func (m *Meta) UnmarshalJSON(data []byte) error {
	type Alias struct {
		CurrentPage int             `json:"current_page"`
		LastPage    int             `json:"last_page"`
		PerPage     json.RawMessage `json:"per_page"`
		Total       int             `json:"total"`
		Seed        string          `json:"seed"`
		SearchQuery json.RawMessage `json:"query"`
	}

	var a Alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	m.CurrentPage = a.CurrentPage
	m.LastPage = a.LastPage
	m.Total = a.Total
	m.Seed = a.Seed
	m.SearchQuery = a.SearchQuery

	var perPageInt int
	if err := json.Unmarshal(a.PerPage, &perPageInt); err == nil {
		m.PerPage = perPageInt
		return nil
	}

	var perPageStr string
	if err := json.Unmarshal(a.PerPage, &perPageStr); err != nil {
		return fmt.Errorf("per_page: unexpected type: %w", err)
	}

	n, err := strconv.Atoi(perPageStr)
	if err != nil {
		return fmt.Errorf("per_page: cannot convert %q to int: %w", perPageStr, err)
	}

	m.PerPage = n
	return nil
}
