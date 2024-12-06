package processors

import (
	"context"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type RecipesRepository interface {
	RecipesAll(ctx context.Context) ([]*domain.Recipe, error)
}

type RecipesLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type recipes struct {
	recipesRepository RecipesRepository
	log               RecipesLogger
}

func NewRecipe(recipesRepository RecipesRepository, log RecipesLogger) *recipes {
	return &recipes{recipesRepository: recipesRepository, log: log}
}

func (res *recipes) RecipesList(ctx context.Context) ([]*domain.Recipe, error) {
	r, err := res.recipesRepository.RecipesAll(ctx)
	if err != nil {
		res.log.Error("recipes list processor error")
		return nil, fmt.Errorf("recipes list processor error: %w", err)
	}
	return r, nil
}

//RecipesAll()
//RecipesToIngredients()
//IngredientsAll()
