package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	openapi "github.com/donskova1ex/mylearningproject/openapi"
)

func main() {
	//logger start
	logger := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	logger.Info("application start")
	slog.SetDefault(logger)

	//logger end
	IngredientAPIService := openapi.NewIngredientAPIService()
	IngredientAPIController := openapi.NewIngredientAPIController(IngredientAPIService)

	RecipeAPIService := openapi.NewRecipeAPIService()
	RecipeAPIController := openapi.NewRecipeAPIController(RecipeAPIService)

	WitchAPIService := openapi.NewWitchAPIService()
	WitchAPIController := openapi.NewWitchAPIController(WitchAPIService)

	router := openapi.NewRouter(IngredientAPIController, RecipeAPIController, WitchAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))

}
