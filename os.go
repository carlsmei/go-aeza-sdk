package aeza_sdk

import (
	"encoding/json"
	"errors"
)

type OS struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Repository string `json:"repository"`
	Group      string `json:"group"`
	Enabled    bool   `json:"enabled"`
}

func (client *Client) GetOSList() ([]OS, error) {
	var res Response

	if _, err := client.restyClient.R().SetResult(&res).Get("os"); err != nil {
		return nil, err
	}

	if res.Error.Slug != "" {
		return nil, errors.New(res.Error.Message)
	}

	var items []OS

	if err := json.Unmarshal(res.Data.Items, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (client *Client) GetOS(id int) (*OS, error) {
	osList, err := client.GetOSList()
	if err != nil {
		return nil, err
	}

	for _, v := range osList {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("OS Not Found")
}
