package app_test

import (
	"log"
	"testing"

	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/api"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedFactApiService struct{ mock.Mock }

func (c *MockedFactApiService) Get() (api.ApiCatFact, error) {
	args := c.Called()
	return args.Get(0).(api.ApiCatFact), args.Error(1)
}

func TestBetterFactService(t *testing.T) {
	mockedApiGetter := new(MockedFactApiService)
	fact := "Cats have four legs"
	mockedApiGetter.On("Get").Return(api.ApiCatFact{Fact: fact, Length: 19}, nil).Once()
	betterFactService := app.NewBetterFactService(mockedApiGetter)
	betterFact, err := betterFactService.Get()
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, fact, betterFact.BestFactEver)
}
