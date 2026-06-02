package gopaper

import (
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
