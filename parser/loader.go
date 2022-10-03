package parser

import (
	"appmagic-entrance/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func LoadJSON(url string) (*models.Ethereum, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("cannot load from url with err: " + err.Error())
	}

	if response.StatusCode/100 != 2 {
		return nil, errors.New("error: not successful response from server")
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("cannot read JSON with err: " + err.Error())
	}

	var result *models.Ethereum
	json.Unmarshal(contents, &result)

	return result, nil
}
