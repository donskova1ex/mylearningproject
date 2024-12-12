package processors

import (
	"context"
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

//go:generate mockgen -destination=./mocks/ingredients_repository.go -package=mocks -mock_names=IngredientsRepository=IngredientsRepository . IngredientsRepository
type IngredientsRepository interface {
	IngredientsAll(ctx context.Context) ([]*domain.Ingredient, error)
	IngredientByUUID(ctx context.Context, uuid string) (*domain.Ingredient, error)
	DeleteIngredientByUUID(ctx context.Context, uuid string) error
	UpdateIngredientByUUID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error)
	CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error)
}

//go:generate mockgen -destination=./mocks/ingredients_logger.go -package=mocks -mock_names=IngredientsLogger=IngredientsLogger . IngredientsLogger
type IngredientsLogger interface { //TODO: добавить во все
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type ingredients struct {
	ingredientsRepository IngredientsRepository
	log                   IngredientsLogger
}

func NewIngredient(ingredientsRepository IngredientsRepository, log IngredientsLogger) *ingredients {
	return &ingredients{ingredientsRepository: ingredientsRepository, log: log}
}

func (i *ingredients) IngredientsList(ctx context.Context) ([]*domain.Ingredient, error) {

	r, err := i.ingredientsRepository.IngredientsAll(ctx)
	if err != nil {
		i.log.Error("it is impossible to get a ingredients list", slog.String("err", err.Error()))
		return nil, err
	}

	return r, nil
}

func (i *ingredients) IngredientByID(ctx context.Context, uuid string) (*domain.Ingredient, error) {
	ing, err := i.ingredientsRepository.IngredientByUUID(ctx, uuid)
	if err != nil {
		i.log.Error("unable to get ingredient by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid)) //TODO:везде логировать
		return nil, err
	}
	return ing, nil
}
func (i *ingredients) DeleteIngredientByID(ctx context.Context, uuid string) error {
	err := i.ingredientsRepository.DeleteIngredientByUUID(ctx, uuid)
	if err != nil {
		i.log.Error("unable to delete ingredient by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return err
	}
	return nil
}

func (i *ingredients) UpdateIngredientByID(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	ing, err := i.ingredientsRepository.UpdateIngredientByUUID(ctx, ingredient)
	if err != nil {
		i.log.Error("unable to update ingredient by uuid")
		return nil, err
	}
	return ing, nil
}

func (i *ingredients) CreateIngredient(ctx context.Context, ingredient *domain.Ingredient) (*domain.Ingredient, error) {
	ing, err := i.ingredientsRepository.CreateIngredient(ctx, ingredient)

	if err != nil {
		i.log.Error("unable to create ingredient",
			slog.String("err", err.Error()))
		return nil, err
	}

	return ing, nil
}
