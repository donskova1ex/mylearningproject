package repositories

import (
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/jmoiron/sqlx"
)

type IngredientsPostgres struct {
	db *sqlx.DB
}

func NewIngredientPostgres(db *sqlx.DB) *IngredientsPostgres {
	return &IngredientsPostgres{db: db}
}

func (i *IngredientsPostgres) CreateIngredient(ingredient domain.Ingredient) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id")
	row := i.db.QueryRow(query, ingredient.Name)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (i *IngredientsPostgres) IngredientsAll() ([]*domain.Ingredient, error) {
	ingredients := []*domain.Ingredient{}
	rows, error := i.db.Queryx("SELECT * FROM Ingredients")
	if error != nil {
		return nil, error
	}
	for rows.Next() {
		error = rows.StructScan(&ingredients)
		if error != nil {
			return nil, error
		}
	}
	return ingredients, nil
}
