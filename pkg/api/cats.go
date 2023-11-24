package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApiCatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type CatFactApi struct {
	apiAddress string
}

func NewFactApi(apiAddress string) CatFactApi {
	return CatFactApi{
		apiAddress: apiAddress,
	}
}

func (c CatFactApi) Get() (ApiCatFact, error) {
	req, err := http.NewRequest(http.MethodGet, c.apiAddress, nil)
	if err != nil {
		return ApiCatFact{}, err
	}

	client := http.DefaultClient
	res, getErr := client.Do(req)
	if getErr != nil {
		return ApiCatFact{}, getErr
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return ApiCatFact{}, readErr
	}

	var catFact ApiCatFact
	parseErr := json.Unmarshal(body, &catFact)
	if parseErr != nil {
		return ApiCatFact{}, parseErr
	}
	return catFact, nil
}
