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
	err := row.Err()
	if err != nil {
		return nil, fmt.Errorf("can not read recipe from db: %w", err)
	}
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("can not scan recipe for id: %w", err)
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

func (r *RecipesPostgres) RecipeByUUID(ctx context.Context, uuid string) (*domain.Recipe, error) {
	recipe := &domain.Recipe{}
	query := "SELECT uuid, id, name, brew_time_seconds FROM recipes WHERE uuid = $1"
	row := r.db.QueryRow(query, uuid)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, fmt.Errorf("recipe with UUID: %s not found: %w", uuid, row.Err())
	}
	if row.Err() != nil {
		return nil, fmt.Errorf("can not read recipe from db: %w", row.Err())
	}
	err := row.Scan(&recipe)
	if err != nil {
		return nil, fmt.Errorf("can not create struct recipe from db: %w", err)
	}
	return recipe, nil
}

func (r *RecipesPostgres) DeleteRecipeByUUID(ctx context.Context, uuid string) error {
	_, err := r.db.Exec("DELETE FROM recipes WHERE uuid = $1", uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("recipe not found: %w", err)
	}
	return nil
}

func (r *RecipesPostgres) UpdateRecipeByUUID(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {
	query := "UPDATE recipes SET name = $1, brew_time_seconds = $2 WHERE uuid = $3"
	_, err := r.db.Exec(query, recipe.Name, recipe.BrewTimeSeconds, recipe.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return recipe, nil
}
