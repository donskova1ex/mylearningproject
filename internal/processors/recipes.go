package processors

import (
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/donskova1ex/mylearningproject/internal/repositories"
	"github.com/jmoiron/sqlx"
)

type RecipesRepository interface {
	NewRecipePostgres(db *sqlx.DB) *repositories.RecipesPostgres
}

type recipes struct {
	recipesRepository RecipesRepository
	log               *slog.Logger
}

func NewRecipes(recipesRepository RecipesRepository, log *slog.Logger) *recipes {
	return &recipes{recipesRepository: recipesRepository, log: log}
}

func (p *recipes) RecipesList() ([]*domain.Recipe, error) {
	db, err := repositories.DBConnection()
	if err != nil {
		return nil, err
	}

	r := p.recipesRepository.NewRecipePostgres(db)

	recipes, err := r.RecipesAll()
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

//RecipesAll()
//RecipesToIngredients()
//IngredientsAll()
