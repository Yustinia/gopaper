package gopaper

type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient creates a new unauthorized client
func NewClient() Client {
	return Client{BaseURL: "https://wallhaven.cc/api/v1"}
}

// NewClientWithKey creates a new authorized client
func NewClientWithKey(apiKey string) Client {
	return Client{APIKey: apiKey, BaseURL: "https://wallhaven.cc/api/v1"}
}
