package aeza_sdk

import (
	"encoding/json"
	"fmt"
)

type Parameters struct {
	ISOUrl            string  `json:"isoUrl"`
	OSID              int     `json:"os"`
	Recipe            *int    `json:"recipe,omitempty"`
	PanelUsername     *string `json:"panelUsername,omitempty"`
	DDoSNotifications *bool   `json:"ddosNotifications,omitempty"`
}

type SecureParameters struct {
	UnSecure bool              `json:"unsecure"`
	Data     map[string]string `json:"data"`
}

type Service struct {
	ID               int              `json:"id"`
	OwnerID          int              `json:"ownerId"`
	ProductID        int              `json:"productId"`
	Name             string           `json:"name"`
	IP               string           `json:"ip"`
	PaymentTerm      string           `json:"paymentTerm"`
	Parameters       Parameters       `json:"parameters"`
	SecureParameters SecureParameters `json:"secureParameters"`
	AutoProlong      bool             `json:"autoProlong"`
	Backups          bool             `json:"backups"`
	Status           string           `json:"status"`
	LastStatus       string           `json:"lastStatus"`
	Product          Product          `json:"product"`
	LocationCode     string           `json:"locationCode"`
	Prices           map[string]Price `json:"prices"`
	CurrentStatus    string           `json:"currentStatus"`
}

func (client *Client) GetServices() []Service {
	var res Response

	client.restyClient.R().SetResult(&res).Get("services")

	var items []Service
	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	return items
}

func (client *Client) GetService(id int) Service {
	var res Response

	client.restyClient.R().SetResult(&res).Get(fmt.Sprintf("services/%d", id))

	if res.Error.Slug != "" {
		panic("err")
	}

	var items []Service
	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	return items[0]
}

type DeleteServiceResponse struct {
	Data  string `json:"data,omitempty"`
	Error *Error `json:"error,omitempty"`
}

func (client *Client) DeleteService(id int) bool {
	var res DeleteServiceResponse

	client.restyClient.R().SetResult(&res).Delete(fmt.Sprintf("services/%d", id))

	if res.Error.Slug != "" {
		panic("err")
	}

	return res.Data == "ok"
}

// func (client *Client) ChangeService(id int) Service {
// 	var res Response

// 	client.restyClient.R().SetResult(&res).Put(fmt.Sprintf("services/%d", id))

// 	if res.Error.Slug != "" {
// 		panic("err")
// 	}

// 	var items []Service
// 	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
// 		panic(err)
// 	}

// 	return items[0]
// }

type ChangeServicePasswordDTO struct {
	Password string `json:"password"`
}

func (client *Client) ChangeServicePassword(id int, password string) Service {
	var res Response

	client.restyClient.R().
		SetResult(&res).
		SetBody(&ChangeServicePasswordDTO{
			Password: password,
		}).
		Post(fmt.Sprintf("services/%d/changePassword", id))

	if res.Error.Slug != "" {
		panic("err")
	}

	var items []Service
	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	return items[0]
}

// https://my.aeza.net/api/services/%d/tasks
// func (client *Client) GetServiceTasks() {
//
// }
