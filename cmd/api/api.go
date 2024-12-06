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
	db, err := repositories.NewPostgresDB(repositories.Config{})
	if err != nil {
		logger.Error("can not postgres db connection", slog.String("error", err.Error()))
		return
	}
	defer db.Close()
	ingredientRepository := repositories.NewIngredientPostgres(db)
	ingProcessor := processors.NewIngredient(ingredientRepository, logger)

	IngredientAPIService := openapi.NewIngredientAPIService(ingProcessor, logger)
	IngredientAPIController := openapi.NewIngredientAPIController(IngredientAPIService)

	RecipeAPIService := openapi.NewRecipeAPIService(, logger)
	RecipeAPIController := openapi.NewRecipeAPIController(RecipeAPIService)

	WitchAPIService := openapi.NewWitchAPIService()
	WitchAPIController := openapi.NewWitchAPIController(WitchAPIService)

	router := openapi.NewRouter(IngredientAPIController, RecipeAPIController, WitchAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))

}
