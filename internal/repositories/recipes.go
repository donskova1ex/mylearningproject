package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RecipesPostgres struct {
	db *sqlx.DB
}

func NewRecipePostgres(db *sqlx.DB) *RecipesPostgres {
	return &RecipesPostgres{db: db}
}

func (r *RecipesPostgres) CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {
	var id uint32
	query := "INSERT INTO recipes (uuid, Name, BrewTimeSeconds) values ($1, $2, $3)RETURNING id"
	newUUID := uuid.NewString()
	row := r.db.QueryRow(query, newUUID, recipe.Name, recipe.BrewTimeSeconds)
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("impossible to create an entity: %w", err)
	}

	newRecipe := &domain.Recipe{
		UUID:            newUUID,
		ID:              id,
		Name:            recipe.Name,
		BrewTimeSeconds: recipe.BrewTimeSeconds,
	}
	return newRecipe, nil
}

func (r *RecipesPostgres) RecipesAll(ctx context.Context) ([]*domain.Recipe, error) {
	recipes := []*domain.Recipe{}

	rows, err := r.db.Queryx("SELECT uuid, id, name, brew_time_seconds FROM recipes")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("empty table: %w", err)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&recipes)
		if err != nil {
			return nil, fmt.Errorf("unable to perform ingredients select: %w", err)
		}
	}
	return recipes, nil
}
