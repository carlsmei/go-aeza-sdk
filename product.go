package aeza_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Configuration struct {
	Max  int    `json:"max"`
	Base int    `json:"base"`
	Slug string `json:"slug"`
	Type string `json:"type"`
	// Prices
}

type Group struct {
	ID    int     `json:"id"`
	Order int     `json:"order"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Role  *string `json:"role"`
}

type Product struct {
	ID           int    `json:"id"`
	Type         string `json:"type"`
	GroupID      int    `json:"groupId"`
	Order        int    `json:"order"`
	InstallPrice int    `json:"installPrice"`
	Name         string `json:"name"`
	// DefaultParameters
	Configuration  []Configuration  `json:"configuration"`
	RawPrices      map[string]int   `json:"rawPrices"`
	IsPrivate      bool             `json:"isPrivate"`
	Group          Group            `json:"group"`
	Prices         map[string]Price `json:"prices"`
	ServiceHandler string           `json:"serviceHandler"`
}

type BuyProductDTO struct {
	AutoProlong bool       `json:"autoProlong"`
	Backups     bool       `json:"backups"`
	Count       int        `json:"count"`
	Method      string     `json:"method"`
	Parameters  Parameters `json:"parameters"`
	Name        string     `json:"name"`
	ProductID   int        `json:"productId"`
	Term        string     `json:"term"`
}

type Transaction struct {
	ID          int             `json:"id"`
	OwnerId     int             `json:"ownerId"`
	Amount      int             `json:"amount"`
	BonusAmount int             `json:"bonusAmount"`
	Mode        string          `json:"mode"`
	Status      string          `json:"status"`
	PerformedAt int             `json:"performedAt"`
	CreatedAt   int             `json:"createdAt"`
	Type        string          `json:"type"`
	OrderInfo   json.RawMessage `json:"orderInfo"`
}

func (client *Client) GetProducts() ([]Product, error) {
	var res Response

	if _, err := client.restyClient.R().SetResult(&res).Get("services/products"); err != nil {
		return nil, err
	}

	if res.Error.Slug != "" {
		return nil, errors.New(res.Error.Message)
	}

	var items []Product

	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (client *Client) GetProduct(id int) (*Product, error) {
	var res Response

	if _, err := client.restyClient.R().SetResult(&res).Get(fmt.Sprintf("services/products/%d", id)); err != nil {
		return nil, err
	}

	if res.Error.Slug != "" {
		return nil, errors.New(res.Error.Message)
	}

	var items []Product

	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("no product found")
	}

	return &items[0], nil
}

func (client *Client) BuyProduct(name string, autoProlong bool, osId int, productId int, term string) (*Service, *Transaction, error) {
	product, _ := client.GetProduct(productId)

	if product.Type != "hicpu" && product.Type != "vps" {
		return nil, nil, errors.New("not implemented (hicpu or vps only)")
	}

	var res Response
	if _, err := client.restyClient.R().
		SetResult(&res).
		SetHeader("Content-Type", "application/json").
		SetBody(BuyProductDTO{
			AutoProlong: autoProlong,
			Backups:     false,
			Count:       1,
			Method:      "balance",
			Name:        name,
			Parameters: Parameters{
				ISOUrl: "",
				OSID:   osId,
			},
			ProductID: productId,
			Term:      term,
		}).
		Post("services/orders"); err != nil {
		return nil, nil, err
	}

	if res.Error.Slug != "" {
		return nil, nil, errors.New(res.Error.Message)
	}

	var items []Service
	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		return nil, nil, err
	}

	service := items[0]
	transaction := res.Data.Transaction

	return &service, transaction, nil
}

// func (product *Product) Buy() {
// 	if product.Type != "hicpu" && product.Type != "vps" {
// 		panic("not implemented")
// 	}

// 	var res Response
// 	product.restyClient.R().SetResult(&res).Get("accounts?current=1&edit=true&extra=1")

// 	var items []Account
// 	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
// 		panic(err)
// 	}

// 	account := items[0]
// 	account.sdkClient = client

// 	return &account, nil
// }
