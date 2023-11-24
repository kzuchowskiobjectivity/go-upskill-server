package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/api"
	"github.com/kzuchowskiobjectivity/go-upskill-server/pkg/app"
	ihttp "github.com/kzuchowskiobjectivity/go-upskill-server/pkg/http"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	configFlag := flag.String("configPath", ".", "Set path to config")
	flag.Parse()
	viper.AddConfigPath(*configFlag)
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
	factApiGetter := api.NewFactApi(factsEndpoint)
	betterFactSvc := app.NewBetterFactService(factApiGetter)
	r := gin.Default()
	handler := ihttp.NewHandler(betterFactSvc)
	ihttp.Routes(&r.RouterGroup, handler)

	portNumber := viper.GetString("port")
	runErr := r.Run(portNumber)
	if runErr != nil {
		log.Fatal(err)
	}
}
