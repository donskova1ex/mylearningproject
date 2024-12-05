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
	query := "INSERT INTO recipes (Name, BrewTimeSeconds, Ingredients) values ($1, $2, $3)RETURNING id" //???
	row := r.db.QueryRow(query, recipe.Name, recipe.BrewTimeSeconds, recipe.Ingredients)                // ???
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("impossible to create an entity: %w", err)
	}
	return id, nil
}

func (r *RecipesPostgres) RecipesAll() ([]*domain.Recipe, error) {
	recipes := []*domain.Recipe{}

	rows, err := r.db.Queryx("SELECT * FROM recipes")
	if err != nil {
		return nil, err //TODO: Обернуть
	}

	for rows.Next() {
		err = rows.Scan(&recipes)
		if err != nil {
			return nil, err //TODO: Обернуть
		}
	}
	return recipes, nil
}
