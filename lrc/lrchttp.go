package lrc

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Client struct {
	http.Client
}

func NewClient() *Client {
	return &Client{
		http.Client{
			Timeout: time.Second * 2,
		},
	}
}

func (c *Client) Get(url string, referer string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:77.0) Gecko/20100101 Firefox/77.0")
	if referer != "" {
		req.Header.Add("Referer", referer)
	}
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func ParseJson(resp *http.Response, jsonRes *map[string]interface{}) error {
	if err := json.NewDecoder(resp.Body).Decode(&jsonRes); err != nil {
		log.Println("Error parsing response into json : ", err)
		return err
	}
	return nil
}
