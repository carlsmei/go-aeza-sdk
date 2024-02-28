package aeza_sdk

import "encoding/json"

type OS struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Repository string `json:"repository"`
	Group      string `json:"group"`
	Enabled    bool   `json:"enabled"`
}

func (client *Client) GetOSList() []OS {
	var res Response

	client.restyClient.R().SetResult(&res).Get("os")

	var items []OS

	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	return items
}

func (client *Client) GetOS(id int) *OS {
	var res Response

	client.restyClient.R().SetResult(&res).Get("os")

	var items []OS

	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		panic(err)
	}

	for _, v := range items {
		if v.ID == id {
			return &v
		}
	}

	panic("not found")
}
