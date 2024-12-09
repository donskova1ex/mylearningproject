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
		return nil, fmt.Errorf("impossible to create an entity: %w", err) //TODO: обертка ошибок через fmt.Errorf
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
		return nil, fmt.Errorf("empty table: %w", err)
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

func (i *IngredientsPostgres) IngredientByID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{}
	query := "SELECT id, name, uuid FROM ingredients WHERE uuid = $1"
	err := i.db.QueryRow(query, uuid).Scan(&ingredient)
	if err != nil {
		return ingredient, nil
	}
	return ingredient, fmt.Errorf("there is no object with this ID: %w", err)
}

func (i *IngredientsPostgres) DeleteIngredientByID(ctx context.Context, uuid string) error {
	_, err := i.db.Exec("DELETE FROM ingredients WHERE uuid = $1", uuid)

	if err != nil {
		return fmt.Errorf("there is no object with this ID: %w", err)
	}
	return nil
}

func (i *IngredientsPostgres) UpdateIngredientByID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	query := "UPDATE ingredients SET name = $1 WHERE uuid = $2"
	_, err := i.db.Exec(query, ingredient.Name, ingredient.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return ingredient, nil
}
