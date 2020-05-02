package remofwd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() Client {
	return Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) FetchDevices() ([]Device, error) {
	req, _ := http.NewRequest("GET", "https://api.nature.global/1/devices", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("API_TOKEN")))
	res, _ := c.httpClient.Do(req)
	data, _ := ioutil.ReadAll(res.Body)

	var result []Device

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
