package repositories

import (
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/jmoiron/sqlx"
)

type RecipesPostgres struct {
	db *sqlx.DB
}

func NewRecipePostgres(db *sqlx.DB) *RecipesPostgres {
	return &RecipesPostgres{db: db}
}

func (r *RecipesPostgres) CreateRecipe(recipe domain.Recipe) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO recipes (Name, BrewTimeSeconds, Ingredients) values ($1, $2, $3)RETURNING id") //???
	row := r.db.QueryRow(query, recipe.Name, recipe.BrewTimeSeconds, recipe.Ingredients)                             // ???
	if error := row.Scan(&id); error != nil {
		return "", error
	}
	return id, nil
}

func (r *RecipesPostgres) RecipesAll() ([]*domain.Recipe, error) {
	recipes := []*domain.Recipe{}

	rows, error := r.db.Queryx("SELECT * FROM recipes")
	if error != nil {
		return nil, error
	}

	for rows.Next() {
		error = rows.Scan(&recipes)
		if error != nil {
			return nil, error
		}
	}
	return recipes, nil
}
