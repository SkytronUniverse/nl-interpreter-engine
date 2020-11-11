package sentiment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nl-interpreter-engine/web"
)

const (
	//URL of sentim-API
	URL = "https://sentim-api.herokuapp.com/api/v1/"
)

//Sentimenter interface defines sentiment interpretation functions
type Sentimenter interface {
	InterpretSentiment(http web.Client) ([]byte, error)
}

//Client struct manages api requests
type Client struct {
	web.HTTPClient
}

//New creates new sentiment client
func New(w web.HTTPClient) *Client {
	return &Client{w}
}

//InterpretSentiment calls API to interpret sentiment
func (c *Client) InterpretSentiment(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	jsonStr := fmt.Sprintf(`{ "text" : "%s" }`, strings.Replace(string(data), "\n", " ", -1))
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
