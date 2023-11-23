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

func NewFactApi(apiAddress string) domain.FactApiService {
	return CatFactApi{
		apiAddress: apiAddress,
	}
}

func (c CatFactApi) Get() (domain.ApiCatFact, error) {
	var catFact domain.ApiCatFact
	req, err := http.NewRequest(http.MethodGet, c.apiAddress, nil)
	if err != nil {
		return catFact, err
	}
	client := http.DefaultClient
	res, getErr := client.Do(req)
	if getErr != nil {
		// more idiomatic
		// should return fresh domain.ApiCatFact{}
		return catFact, getErr
	}

	// from documentation:
	// " The http Client and Transport guarantee that Body is always
	// non-nil, even on responses without a body or responses with
	// a zero-length body."
	// so no need to nil check, but error check must be done before close
	// https://medium.com/@KeithAlpichi/go-gotcha-closing-a-nil-http-response-body-with-defer-9b7a3eb30e8c
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

type BetterFactService struct {
	api domain.FactApiService
}

func NewBetterFactService(api domain.FactApiService) domain.BetterFactService {
	return BetterFactService{
		api: api,
	}
}

func (svc BetterFactService) Get() (domain.BetterCatFact, error) {
	fact, err := svc.api.Get()
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
