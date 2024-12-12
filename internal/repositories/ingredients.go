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

func (i *IngredientsPostgres) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	var id uint32
	var newUUID string
	var uuidNotUsed = false

	query := "INSERT INTO ingredients (name, uuid) values ($1, $2) RETURNING id"

	//перебор uuid
	for uuidNotUsed != true {
		newUUID = uuid.NewString()
		checkingRow := i.db.QueryRowContext(ctx, "SELECT uuid FROM ingredients WHERE uuid = $1", newUUID)
		if err := checkingRow.Scan(&newUUID); err == nil {
			uuidNotUsed = true
		}
	}
	//перебор uuid

	row := i.db.QueryRow(query, ingredient.Name, newUUID)
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

func (i *IngredientsPostgres) IngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ingredient := &domain.Ingredient{}
	query := "SELECT id, name, uuid FROM ingredients WHERE uuid = $1"
	row := i.db.QueryRow(query, uuid)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, fmt.Errorf("ingredient with UUID: %s not found: %w", uuid, row.Err())
	}
	if row.Err() != nil {
		return nil, fmt.Errorf("can not read ingredient from db: %w", row.Err())
	}
	err := row.Scan(&ingredient)
	if err != nil {
		return nil, fmt.Errorf("can not create struct ingredient from db: %w", err)
	}
	return ingredient, nil
}

func (i *IngredientsPostgres) DeleteIngredientByUUID(ctx context.Context, uuid string) error {
	_, err := i.db.Exec("DELETE FROM ingredients WHERE uuid = $1", uuid)

	if err != nil {
		return fmt.Errorf("there is no object with this ID: %w", err)
	}
	return nil
}

func (i *IngredientsPostgres) UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	query := "UPDATE ingredients SET name = $1 WHERE uuid = $2"
	_, err := i.db.Exec(query, ingredient.Name, ingredient.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return ingredient, nil
}
