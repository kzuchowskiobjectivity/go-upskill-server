package app

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/domain"
)

type CatFactApi struct {
	apiAddress string
}

func NewFactApiGetter(apiAddress string, client http.Client) domain.BetterFactService {
	return CatFactApi{
		apiAddress: apiAddress,
	}
}

func (c CatFactApi) Get() (domain.BetterCatFact, error) {
	fact, err := c.GetApi()
	var betterFact domain.BetterCatFact
	if err != nil {
		return betterFact, err
	}
	betterFact = domain.BetterCatFact{
		BestFactEver:  fact.Fact,
		UnixTimestamp: time.Now().Unix(),
	}
	return betterFact, nil
}

func (c CatFactApi) GetApi() (domain.ApiCatFact, error) {
	var catFact domain.ApiCatFact
	req, err := http.NewRequest(http.MethodGet, c.apiAddress, nil)
	if err != nil {
		return catFact, err
	}
	client := http.DefaultClient
	res, getErr := client.Do(req)
	if getErr != nil {
		return catFact, getErr
	}
	if res.Body != nil {
		defer res.Body.Close()
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
