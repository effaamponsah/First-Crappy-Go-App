package main

import (
	"food_app/food_app/models"
	"food_app/food_app/restapi"
	"food_app/food_app/restapi/operations"
	"food_app/food_app/restapi/operations/foods"
	"food_app/food_app/restapi/operations/sample_description"

	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

var dummyFoodSource = make(map[int64]*models.Food)

func allItems(since int64, limit int32) (result []*models.Food) {
	result = make([]*models.Food, 0)
	for id, food := range dummyFoodSource {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, food)
		}
	}
	return
}

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

	// implement find all foods
	api.FoodsGetFoodsHandler = foods.GetFoodsHandlerFunc(
		func(parameters foods.GetFoodsParams) middleware.Responder {
			mergedParams := foods.NewGetFoodsParams()
			mergedParams.Since = swag.Int64(0)
			if parameters.Since != nil {
				mergedParams.Since = parameters.Since
			}
			if parameters.Limit != nil {
				mergedParams.Limit = parameters.Limit
			}
			return foods.NewGetFoodsOK().WithPayload(allItems(*mergedParams.Since, *mergedParams.Limit))
		})
	// serve
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
