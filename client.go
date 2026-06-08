package gopaper

type Client struct {
	APIKey  string
	BaseURL string
}

// NewClient returns a client for unauthenticated API requests
func NewClient() Client {
	return Client{BaseURL: "https://wallhaven.cc/api/v1"}
}

// NewClientWithKey returns a client authenticated with the given API key
// Endpoints like fetching NSFW tagged wallpapers require authentication
func NewClientWithKey(apiKey string) Client {
	return Client{APIKey: apiKey, BaseURL: "https://wallhaven.cc/api/v1"}
}
