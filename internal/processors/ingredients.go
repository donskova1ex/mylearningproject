package processors

import (
	"context"
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type IngredientsRepository interface {
	IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error)
}

type ingredients struct {
	ingredientsRepository IngredientsRepository
	log                   *slog.Logger
}

func NewIngredient(ingredientsRepository IngredientsRepository, log *slog.Logger) *ingredients {
	return &ingredients{ingredientsRepository: ingredientsRepository, log: log}
}

func (i *ingredients) IngredientsList(ctx context.Context) ([]*domain.Ingredient, error) {

	r, err := i.ingredientsRepository.IngredientsAll(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
