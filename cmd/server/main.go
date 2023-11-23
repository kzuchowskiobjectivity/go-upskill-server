package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/app"
	ihttp "github.com/kzuchowskiobjectivity/go-upskill-server/pkg/http"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// hardcoded path - bad idea
	// when binary file builded, config must always be located 2 directory above
	// unintuitive approach
	// also go run path/to/main works only when run from same directory as main
	// use execution argument as path or assume that binary and config file are in the same dir
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
	factApiGetter := app.NewFactApi(factsEndpoint)
	betterFactSvc := app.NewBetterFactService(factApiGetter)
	r := gin.Default()
	handler := ihttp.NewHandler(betterFactSvc)
	ihttp.Routes(&r.RouterGroup, handler)

	portNumber := viper.GetString("port")
	r.Run(portNumber)
}
