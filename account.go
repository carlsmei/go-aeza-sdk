package aeza_sdk

import (
	"encoding/json"
)

type Account struct {
	ID              int     `json:"id"`
	EMail           string  `json:"email"`
	Balance         float64 `json:"balance"`
	BonusBalance    float64 `json:"bonusBalance"`
	BonusUsed       float64 `json:"bonusUsed"`
	BonusState      string  `json:"bonusState"`
	CreatedAt       int64   `json:"createdAt"`
	IsSupportLocked bool    `json:"isSupportLocked"`
	IsOnline        bool    `json:"isOnline"`
}

type Limit struct {
	ID        int
	Name      string
	Groups    []int
	Grades    map[string]int
	Available int
	Used      int
}

func (client *Client) GetAccount() (*Account, error) {
	var res Response
	client.restyClient.R().SetResult(&res).Get("accounts?current=1&edit=true&extra=1")

	var items []Account
	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	account := items[0]

	return &account, nil
}

func (client *Client) GetAccountLimits() []Limit {
	var res Response

	client.restyClient.R().SetResult(&res).Get("services/limits")

	var items []Limit

	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	return items
}
