package aeza_sdk

const Version = "1.0.0"

func New(apiKey string) *Client {
	return createClient(apiKey)
}
