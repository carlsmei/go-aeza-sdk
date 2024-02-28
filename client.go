package aeza_sdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type Client struct {
	BaseURL     string
	apiKey      string
	restyClient *resty.Client
}

func createClient(api_key string) *Client {
	resty := resty.New()

	resty.SetBaseURL("https://my.aeza.net/api/")
	resty.SetHeader("X-API-Key", api_key)
	resty.SetHeader("User-Agent", fmt.Sprintf("go-aeza-sdk/%s (github.com/carlsmei/go-aeza-sdk)", Version))

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	resty.JSONMarshal = json.Marshal
	resty.JSONUnmarshal = json.Unmarshal

	return &Client{
		BaseURL:     "https://my.aeza.net/api/",
		restyClient: resty,
		apiKey:      api_key,
	}
}

// func
