package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type FactApiService interface {
	Get() (ApiCatFact, error)
}

type ApiCatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type CatFactApi struct {
	apiAddress string
}

func NewFactApi(apiAddress string) FactApiService {
	return CatFactApi{
		apiAddress: apiAddress,
	}
}

func (c CatFactApi) Get() (ApiCatFact, error) {
	var catFact ApiCatFact
	req, err := http.NewRequest(http.MethodGet, c.apiAddress, nil)
	if err != nil {
		return catFact, err
	}
	client := http.DefaultClient
	res, getErr := client.Do(req)
	if getErr != nil {
		return catFact, getErr
	}
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return catFact, readErr
	}
	parseErr := json.Unmarshal(body, &catFact)
	if parseErr != nil {
		return catFact, parseErr
	}
	return catFact, nil
}
