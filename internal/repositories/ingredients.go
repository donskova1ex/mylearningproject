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

type IngredientsPostgres struct {
	db *sqlx.DB
}

func NewIngredientPostgres(db *sqlx.DB) *IngredientsPostgres {
	return &IngredientsPostgres{db: db}
}

func (i *IngredientsPostgres) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) { //TODO: перенести в рецепты и другие таблицы аналогично
	var id uint32
	query := "INSERT INTO ingredients (name, uuid) values ($1, $2) RETURNING id"
	newUUID := uuid.NewString()
	row := i.db.QueryRow(query, ingredient.Name, newUUID)
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("impossible to create an entity: %w", err) //TODO: обертка ошибок через fmt.Errorf (пример в телеге)
	}
	newIngr := &domain.Ingredient{
		Name: ingredient.Name,
		UUID: newUUID,
		ID:   id,
	}
	return newIngr, nil //TODO: возврат *domain.Ingredient
}

func (i *IngredientsPostgres) IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error) {
	ingredients := []*domain.Ingredient{}
	rows, err := i.db.Queryx("SELECT uuid, id, name FROM ingredients")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("empty answer: %w", err)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&ingredients)
		if err != nil {
			return nil, fmt.Errorf("unable to perform ingredients select: %w", err)
		}
	}
	return ingredients, nil
}
