package gopaper

type Client struct {
	APIKey  string
	BaseURL string
}

func NewClient(apiKey string) Client {
	return Client{APIKey: apiKey, BaseURL: "https://wallhaven.cc/api/v1"}
}
