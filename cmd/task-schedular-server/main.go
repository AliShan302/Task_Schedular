package main

import (
	"strconv"

	loads "github.com/go-openapi/loads"
	"github.com/spf13/viper"

	runtime "github.com/AliShan302/Task_Schedular"
	"github.com/AliShan302/Task_Schedular/config"
	"github.com/AliShan302/Task_Schedular/gen/restapi"
	"github.com/AliShan302/Task_Schedular/handlers"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	rt, err := runtime.NewRuntime()
	if err != nil {
		panic(err)
	}

	api := handlers.NewHandler(rt, swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString(config.ServerHost)
	server.Port, err = strconv.Atoi(viper.GetString(config.ServerPort))
	if err != nil {
		panic(err)
	}
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
