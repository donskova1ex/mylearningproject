package processors

import (
	"context"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

//go:generate mockgen -destination=./mocks/ingredients_repository.go -package=mocks -mock_names=IngredientsRepository=IngredientsRepository . IngredientsRepository
type IngredientsRepository interface {
	IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error)
}

//go:generate mockgen -destination=./mocks/ingredients_logger.go -package=mocks -mock_names=IngredientsLogger=IngredientsLogger . IngredientsLogger
type IngredientsLogger interface { //TODO: добавить во все
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type ingredients struct {
	ingredientsRepository IngredientsRepository
	log                   IngredientsLogger
}

func NewIngredient(ingredientsRepository IngredientsRepository, log IngredientsLogger) *ingredients {
	return &ingredients{ingredientsRepository: ingredientsRepository, log: log}
}

func (i *ingredients) IngredientsList(ctx context.Context) ([]*domain.Ingredient, error) {

	r, err := i.ingredientsRepository.IngredientsAll(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
