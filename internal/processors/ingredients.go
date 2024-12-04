package processors

import (
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/donskova1ex/mylearningproject/internal/repositories"
	"github.com/jmoiron/sqlx"
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
	db, error := repositories.DBConnection()
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
