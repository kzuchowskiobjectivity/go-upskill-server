package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/app"
	ihttp "github.com/kzuchowskiobjectivity/go-upskill-server/pkg/http"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../.")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	logPath := viper.GetString("logPath")
	file, fileErr := os.Create(logPath)
	if fileErr != nil {
		log.Fatal(fileErr)
		return
	}
	gin.DefaultWriter = file

	factsEndpoint := viper.GetString("factEndpoint")
	factApiGetter := app.NewFactApiGetter(factsEndpoint, http.Client{})

	r := gin.Default()
	handler := ihttp.NewHandler(factApiGetter)
	ihttp.Routes(&r.RouterGroup, handler)

	portNumber := viper.GetString("port")
	r.Run(portNumber)
}
