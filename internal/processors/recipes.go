package processors

import (
	"context"
	"errors"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type RecipesRepository interface {
	RecipesAll(ctx context.Context) ([]*domain.Recipe, error)
	CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error)
	RecipeByID(ctx context.Context, uuid string) (*domain.Recipe, error)
	DeleteRecipeByID(ctx context.Context, uuid string) error
	UpdateRecipeByID(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error)
}

type RecipesLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type recipes struct {
	recipesRepository RecipesRepository
	log               RecipesLogger
}

func NewRecipe(recipesRepository RecipesRepository, log RecipesLogger) *recipes {
	return &recipes{recipesRepository: recipesRepository, log: log}
}

func (rec *recipes) RecipesList(ctx context.Context) ([]*domain.Recipe, error) {
	r, err := rec.recipesRepository.RecipesAll(ctx)
	if err != nil {
		rec.log.Error("recipes list getting error: %w", err) // так или не так
		return nil, fmt.Errorf("recipes list getting error: %w", err)
	}
	return r, nil
}

func (rec *recipes) CreateRecipe(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {
	r, err := rec.recipesRepository.CreateRecipe(ctx, recipe)
	if err != nil {
		rec.log.Error("create recipe error: %w", err)
		return nil, fmt.Errorf("create recipe error: %w", err)
	}
	return r, nil
}
func (rec *recipes) RecipeByID(ctx context.Context, uuid string) (*domain.Recipe, error) {
	r, err := rec.RecipeByID(ctx, uuid)
	if err != nil {
		rec.log.Error("recipe by ID getting error: %w", err)
		return nil, err
	}
	return r, nil
}

func (rec *recipes) DeleteRecipeByID(ctx context.Context, uuid string) error {
	err := rec.recipesRepository.DeleteRecipeByID(ctx, uuid)
	if err != nil {
		return err
	}
	rec.log.Error("recipe by ID deleting error", errors.New("deleting by ID error"))
	return nil
}

func (rec *recipes) UpdateRecipeByID(ctx context.Context, recipe *domain.Recipe) (*domain.Recipe, error) {

	r, err := rec.UpdateRecipeByID(ctx, recipe)
	if err != nil {
		rec.log.Error("recipe by ID updating error: %w", err)
		return nil, err
	}
	return r, err
}
