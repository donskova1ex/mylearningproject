package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/donskova1ex/mylearningproject/internal/processors"
	"github.com/donskova1ex/mylearningproject/internal/repositories"
	openapi "github.com/donskova1ex/mylearningproject/openapi"
)

func main() {
	logJSONHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(logJSONHandler)
	logger.Info("application started")
	slog.SetDefault(logger)

	pgDSN := os.Getenv("POSTGRES_DSN")
	if pgDSN == "" {
		logger.Error("empty POSTGRES_DSN")
		os.Exit(1)
	}

	apiPort := os.Getenv("API_PORT")
	if pgDSN == "" {
		logger.Error("empty API_PORT")
		os.Exit(1)
	}

	db, err := repositories.NewPostgresDB(pgDSN)
	if err != nil {
		logger.Error("can not create postgres db connection", slog.String("error", err.Error()))
		return
	}
	defer db.Close()

	ingredientRepository := repositories.NewIngredientPostgres(db)
	ingProcessor := processors.NewIngredient(ingredientRepository, logger)
	IngredientAPIService := openapi.NewIngredientAPIService(ingProcessor, logger)
	IngredientAPIController := openapi.NewIngredientAPIController(IngredientAPIService)

	recipeRepository := repositories.NewRecipePostgres(db)
	recipeProcessor := processors.NewRecipe(recipeRepository, logger)
	RecipeAPIService := openapi.NewRecipeAPIService(recipeProcessor, logger)
	RecipeAPIController := openapi.NewRecipeAPIController(RecipeAPIService)

	witchRepository := repositories.NewWitchesPostgres(db)
	witchProcessor := processors.NewWitch(witchRepository, logger)
	WitchAPIService := openapi.NewWitchAPIService(witchProcessor, logger)
	WitchAPIController := openapi.NewWitchAPIController(WitchAPIService)

	router := openapi.NewRouter(IngredientAPIController, RecipeAPIController, WitchAPIController)

	httpServer := http.Server{
		Addr:     ":" + apiPort,
		ErrorLog: slog.NewLogLogger(logJSONHandler, slog.LevelError),
		Handler:  router,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		logger.Error("failed to start server", slog.String("err", err.Error()))
	}
}
