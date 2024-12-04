package processors

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/donskova1ex/mylearningproject/internal/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type IngredientsRepository interface {
	NewIngredientPostgres(db *sqlx.DB) *repositories.IngredientsPostgres
}

type ingredients struct {
	ingredientsRepository IngredientsRepository
	log                   *slog.Logger
}

func NewIngredient(ingredientsRepository IngredientsRepository, log *slog.Logger) *ingredients {
	return &ingredients{ingredientsRepository: ingredientsRepository, log: log}
}

func (i *ingredients) IngredientsList() ([]*domain.Ingredient, error) {
	db, error := dbConnection()
	if error != nil {
		return nil, error
	}

	r := i.ingredientsRepository.NewIngredientPostgres(db)
	ingredients, error := r.IngredientsAll()

	if error != nil {
		return nil, error
	}
	return ingredients, nil
}

func dbConnection() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf("unable to read env")
	}
	db, err := repositories.NewPostgresDB(repositories.Config{
		Host:     "localhost",
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	})
	if err != nil {
		fmt.Errorf("failed to initialize db")
	}
	return db, nil
}
