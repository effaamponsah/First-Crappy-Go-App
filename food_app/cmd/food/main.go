package main

import (
	"food_app/food_app/restapi"
	"food_app/food_app/restapi/operations"
	"food_app/food_app/restapi/operations/sample_description"

	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	// load embedded swagger spec from a JSON file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// creates a new service and pass the specfile inside
	api := operations.NewFoodAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// set port
	server.Port = 3000

	// Implement base route
	api.SampleDescriptionShowGreetingsToTheAPIHandler = sample_description.ShowGreetingsToTheAPIHandlerFunc(
		func(param sample_description.ShowGreetingsToTheAPIParams) middleware.Responder {
			return sample_description.NewShowGreetingsToTheAPIOK()
		})

	// serve
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
