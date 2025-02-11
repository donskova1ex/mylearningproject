package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"

	"github.com/donskova1ex/mylearningproject/internal"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/google/uuid"
)

func (r *Repository) createIngredients(ctx context.Context, tx *sql.Tx, ingredients []*domain.Ingredient) ([]*domain.Ingredient, error) {

	return nil, nil
}

func (r *Repository) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	var id uint32
	var pqErr *pq.Error
	//TODO: проверить, что возвращает, по идее должен ИД
	query := `INSERT INTO ingredients (name, uuid) values ($1, $2) on conflict on constraint ingredients_name_key do nothing RETURNING id`
	newUUID := uuid.NewString()
	row := r.db.QueryRowContext(ctx, query, ingredient.Name, newUUID)
	err := row.Err()
	if err != nil {
		if errors.As(err, &pqErr) {
			if pqErr.Constraint == "ingredients_uuid_key" {
				return nil, fmt.Errorf("this uuid  is already in use: %w", err)
			}
			if pqErr.Constraint == "ingredients_name_key" {
				return nil, nil
			}
		}
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

func (r *Repository) IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error) {
	var ingredients []*domain.Ingredient
	err := r.db.SelectContext(ctx, &ingredients, "SELECT uuid, id, name FROM ingredients")
	if errors.Is(err, sql.ErrNoRows) {
		return ingredients, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", internal.ErrReadRows)
	}
	return ingredients, nil
}

func (r *Repository) IngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{}
	query := "SELECT id, name, uuid FROM ingredients WHERE uuid = $1"
	err := r.db.GetContext(ctx, ingredient, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrNotFound, uuid)
	}

	if err != nil {
		return nil, fmt.Errorf("%w with uuid [%s]", internal.ErrReadRows, uuid)
	}
	return ingredient, nil
}

func (r *Repository) DeleteIngredientByUUID(ctx context.Context, uuid string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM ingredients WHERE uuid = $1", uuid)
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

func (r *Repository) UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	query := "UPDATE ingredients SET name = $1 WHERE uuid = $2"
	_, err := r.db.ExecContext(ctx, query, ingredient.Name, ingredient.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return ingredient, nil
}
