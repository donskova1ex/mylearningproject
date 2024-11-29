package main

import (
	"log"
	"net/http"

	openapi "github.com/donskova1ex/mylearningproject/openapi"
)

func main() {
	log.Printf("Server started")

	IngredientAPIService := openapi.NewIngredientAPIService()
	IngredientAPIController := openapi.NewIngredientAPIController(IngredientAPIService)

	RecipeAPIService := openapi.NewRecipeAPIService()
	RecipeAPIController := openapi.NewRecipeAPIController(RecipeAPIService)

	WitchAPIService := openapi.NewWitchAPIService()
	WitchAPIController := openapi.NewWitchAPIController(WitchAPIService)

	router := openapi.NewRouter(IngredientAPIController, RecipeAPIController, WitchAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
