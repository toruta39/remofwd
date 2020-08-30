package remofwd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
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
	logrus.Debugf("Authorization: %s", fmt.Sprintf("Bearer %s", os.Getenv("API_TOKEN")))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("API_TOKEN")))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result []Device

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
