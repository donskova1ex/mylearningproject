package repositories

import (
	"context"

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

func (i *IngredientsPostgres) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (uint32, error) { //TODO: возсврат *domain.Ingredient
	var id uint32
	query := "INSERT INTO ingredients (name, uuid) values ($1, $2) RETURNING id"
	row := i.db.QueryRow(query, ingredient.Name, uuid.NewString())
	if err := row.Scan(&id); err != nil {
		return 0, err //TODO: обертка ошибок через fmt.Errorf (пример в телеге)
	}
	return id, nil //TODO: возврат *domain.Ingredient
}

func (i *IngredientsPostgres) IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error) {
	ingredients := []*domain.Ingredient{}
	rows, err := i.db.Queryx("SELECT uuid, id, name FROM ingredients")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&ingredients)
		if err != nil {
			return nil, err
		}
	}
	return ingredients, nil
}
