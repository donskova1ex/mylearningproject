package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/donskova1ex/mylearningproject/internal/processors"
	"github.com/donskova1ex/mylearningproject/internal/repositories"
	openapi "github.com/donskova1ex/mylearningproject/openapi"
)

func main() {
	//logger start
	logger := slog.New(slog.NewJSONHandler(os.Stdin, nil))
	logger.Info("application start")
	slog.SetDefault(logger)
	//logger end
	db := repositories.NewPostgresDB()
	ingredientRepository := repositories.NewIngredientPostgres(db)
	ingProcessor := processors.NewIngredient(ingredientRepository, logger)

	IngredientAPIService := openapi.NewIngredientAPIService(ingProcessor)
	IngredientAPIController := openapi.NewIngredientAPIController(IngredientAPIService)

	RecipeAPIService := openapi.NewRecipeAPIService()
	RecipeAPIController := openapi.NewRecipeAPIController(RecipeAPIService)

	WitchAPIService := openapi.NewWitchAPIService()
	WitchAPIController := openapi.NewWitchAPIController(WitchAPIService)

	router := openapi.NewRouter(IngredientAPIController, RecipeAPIController, WitchAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))

}
