package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/betterfact", handler.GetFact)
}
