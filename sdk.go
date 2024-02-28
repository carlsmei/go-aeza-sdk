package aeza_sdk

const Version = "0.0.1"

func New(apiKey string) *Client {
	return createClient(apiKey)
}
