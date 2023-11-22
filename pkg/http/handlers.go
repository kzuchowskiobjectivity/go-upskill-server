package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/domain"
)

type Handler struct {
	factGetter domain.BetterFactService
}

func NewHandler(factGetter domain.BetterFactService) *Handler {
	return &Handler{
		factGetter: factGetter,
	}
}

func (h *Handler) GetFact(c *gin.Context) {
	fact, err := h.factGetter.Get()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, fact)
}
