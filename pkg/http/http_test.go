package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/domain"
	ihttp "github.com/kzuchowskiobjectivity/go-upskill-server/pkg/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedBetterFactService struct{ mock.Mock }

func (c *MockedBetterFactService) Get() (domain.BetterCatFact, error) {
	args := c.Called()
	return args.Get(0).(domain.BetterCatFact), args.Error(1)
}

func TestGetPet(t *testing.T) {
	mockedApiGetter := new(MockedBetterFactService)
	handler := ihttp.NewHandler(mockedApiGetter)

	testCases := []struct {
		Name         string
		Mock         *mock.Call
		ExpectedCode int
	}{
		{
			"Test OK",
			mockedApiGetter.On("Get").Return(domain.BetterCatFact{BestFactEver: "Cats have four legs"}, nil).Once(),
			http.StatusOK,
		},
		{
			"Test error",
			mockedApiGetter.On("Get").Return(domain.BetterCatFact{}, errors.New("Network error")).Once(),
			http.StatusInternalServerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			handler.GetFact(c)
			assert.Equal(t, testCase.ExpectedCode, w.Code)
		})
	}
}
