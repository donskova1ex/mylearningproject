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

type IngredientsPostgres struct {
	db *sqlx.DB
}

func NewIngredientPostgres(db *sqlx.DB) *IngredientsPostgres {
	return &IngredientsPostgres{db: db}
}

func (i *IngredientsPostgres) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	var id uint32

	query := "INSERT INTO ingredients (name, uuid) values ($1, $2) RETURNING id"
	//TODO: проверка на дубли
	newUUID := uuid.NewString()
	row := i.db.QueryRowContext(ctx, query, ingredient.Name, newUUID)
	err := row.Err()
	if err != nil {
		return nil, fmt.Errorf("can not read ingredient from db: %w", err)
	}
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("impossible to create an entity: %w", err)
	}
	newIngr := &domain.Ingredient{
		Name: ingredient.Name,
		UUID: newUUID,
		ID:   id,
	}
	return newIngr, nil
}

func (i *IngredientsPostgres) IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error) {
	var ingredients []*domain.Ingredient
	err := i.db.Select(&ingredients, "SELECT uuid, id, name FROM ingredients")
	if errors.Is(err, sql.ErrNoRows) {
		return ingredients, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", internal.ErrReadRows)
	}
	return ingredients, nil
}

func (i *IngredientsPostgres) IngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{}
	query := "SELECT id, name, uuid FROM ingredients WHERE uuid = $1"
	err := i.db.GetContext(ctx, ingredient, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}

	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}
	return ingredient, nil
}

func (i *IngredientsPostgres) DeleteIngredientByUUID(ctx context.Context, uuid string) error {
	result, err := i.db.ExecContext(ctx, "DELETE FROM ingredients WHERE uuid = $1", uuid)
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrGetByUUID, uuid)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}

	if rows == 0 {
		return fmt.Errorf("%w with uuid [%s]", internal.ErrNotDelete, uuid)
	}
	return nil
}

func (i *IngredientsPostgres) UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	query := "UPDATE ingredients SET name = $1 WHERE uuid = $2"
	_, err := i.db.ExecContext(ctx, query, ingredient.Name, ingredient.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return ingredient, nil
}
