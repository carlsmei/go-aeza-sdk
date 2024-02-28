package aeza_sdk

import "encoding/json"

type Error struct {
	Slug    string            `json:"slug"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type Response struct {
	Data struct {
		SelectorMode *string         `json:"selectorMode,omitempty"`
		Items        json.RawMessage `json:"items"`
		TotalItems   *int            `json:"total,omitempty"`
		Transaction  *Transaction    `json:"transaction,omitempty"`
	} `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

type Price struct {
	Value           int    `json:"value"`
	Suffix          string `json:"suffix"`
	DefaultCurrency bool   `json:"defaultCurrency"`
	Slug            string `json:"slug"`
}
