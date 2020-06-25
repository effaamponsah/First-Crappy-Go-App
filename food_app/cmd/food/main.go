package main

import (
	"flag"
	"food_app/food_app/restapi"
	"food_app/food_app/restapi/operations"
	"log"

	"github.com/go-openapi/loads"
)

// var port = flag.Int("port", 5000, "The server port")

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

	// parse flags
	flag.Parse()
	// set port
	server.Port = 3000

	// serve
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
