package app

import (
	"time"

	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/api"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/domain"
)

type BetterFactService struct {
	api api.FactApiService
}

func NewBetterFactService(api api.FactApiService) BetterFactService {
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
