package gopaper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func doRequest[respType any](url string) (respType, error) {
	var result respType

	resp, err := http.Get(url)
	if err != nil {
		return result, fmt.Errorf("something went wrong: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("failed to decode: %w", err)
	}

	return result, nil
}
