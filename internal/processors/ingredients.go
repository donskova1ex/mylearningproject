package processors

import (
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type IngredientsRepository interface {
}

type ingredients struct {
	ingredientsRepository IngredientsRepository
	log                   *slog.Logger
}

func NewIngredient(ingredientsRepository IngredientsRepository, log *slog.Logger) *ingredients {
	return &ingredients{ingredientsRepository: ingredientsRepository, log: log}
}

func (i *ingredients) IngredientsList() ([]*domain.Ingredient, error) {

}
