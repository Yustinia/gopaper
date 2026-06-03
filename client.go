package gopaper

type Client struct {
	APIKey  string
	BaseURL string
}

func NewClient() Client {
	return Client{BaseURL: "https://wallhaven.cc/api/v1"}
}

func NewClientWithKey(apiKey string) Client {
	return Client{APIKey: apiKey, BaseURL: "https://wallhaven.cc/api/v1"}
}
