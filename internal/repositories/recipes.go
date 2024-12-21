package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal"
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
	var recipes []*domain.Recipe

	err := r.db.SelectContext(ctx, &recipes, "SELECT uuid, id, name, brew_time_seconds FROM recipes")
	if errors.Is(err, sql.ErrNoRows) {
		return recipes, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", internal.ErrReadRows)
	}
	return recipes, nil
}

func (r *RecipesPostgres) RecipeByUUID(ctx context.Context, uuid string) (*domain.Recipe, error) {
	recipe := &domain.Recipe{}
	query := "SELECT uuid, id, name, brew_time_seconds FROM recipes WHERE uuid = $1"
	err := r.db.GetContext(ctx, recipe, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid) // TODO: придумать как сделать прим. так же
	}
	if err != nil {
		return nil, fmt.Errorf("%w by uuid: [%s]", internal.ErrReadRows, uuid)
	}
	return recipe, nil
}

func (r *RecipesPostgres) DeleteRecipeByUUID(ctx context.Context, uuid string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM recipes WHERE uuid = $1", uuid)

	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrGetByUUID, uuid)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}

	if rows != 1 {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotDelete, uuid)
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
