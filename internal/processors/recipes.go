package processors

import (
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type RecipesRepository interface {
}

type recipes struct {
	recipesRepository RecipesRepository
	log               *slog.Logger
}

func NewRecipes(recipesRepository RecipesRepository, log *slog.Logger) *recipes {
	return &recipes{recipesRepository: recipesRepository, log: log}
}

func (p *recipes) RecipesList() ([]*domain.Recipe, error) {

}

//RecipesAll()
//RecipesToIngredients()
//IngredientsAll()
