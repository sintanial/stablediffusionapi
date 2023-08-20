package stablediffusionapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const apiUrl = "https://stablediffusionapi.com"

type RequestBase struct {
	Key string `json:"key,omitempty"`
}

type Status = string

const StatusError Status = "error"
const StatusSuccess Status = "success"

type ResponseBase struct {
	Status  Status              `json:"status"`
	Messege map[string][]string `json:"messege"`
}

const DefaultWidth = 512
const DefaultHeight = 512
const DefaultSamples = 4

type RequiredImageFields struct {
	//Max Width: Width: 1024x1024.
	Width int `json:"width"`
	//Max Height: Width: 1024x1024.
	Height int `json:"height"`
	//Number of images to be returned in response. The maximum value is 4.
	Samples int `json:"samples"`
}

func DoPost(client *http.Client, url string, request any, response any) error {
	var reqbody bytes.Buffer
	if err := json.NewEncoder(&reqbody).Encode(request); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl+url, &reqbody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New("invalid status code " + strconv.Itoa(resp.StatusCode) + ", with body: " + string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, response); err != nil {
		return err
	}

	return nil
}

func SetApiKeyIfNeeded(req *RequestBase, apiKey string) {
	if req.Key == "" {
		req.Key = apiKey
	}
}
func SetDefaultRequiredImageFieldsIfNeeded(req *RequiredImageFields) {
	if req.Width == 0 {
		req.Width = DefaultWidth
	}
	if req.Height == 0 {
		req.Height = DefaultHeight
	}
	if req.Samples == 0 {
		req.Samples = DefaultSamples
	}
}
