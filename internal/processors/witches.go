package processors

import (
	"context"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type WitchesRepository interface {
	CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error)
	WitchesAll(ctx context.Context) ([]*domain.Witch, error)
	WitchByID(ctx context.Context, uuid string) (*domain.Witch, error)
	DeleteWitchByID(ctx context.Context, uuid string) error
	UpdateWitchByID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error)
}

type WitchesLogger interface {
	Error(msg string, args ...any)
	Info(msg string, args ...any)
}

type witches struct {
	witchesRepository WitchesRepository
	log               WitchesLogger
}

func NewWitch(witchesRepository WitchesRepository, log WitchesLogger) *witches {
	return &witches{witchesRepository: witchesRepository, log: log}
}

func (wtch *witches) WitchesList(ctx context.Context) ([]*domain.Witch, error) {
	w, err := wtch.witchesRepository.WitchesAll(ctx)
	if err != nil {
		wtch.log.Error("witches list processor error")
		return nil, fmt.Errorf("witches list processor error: %w", err)
	}
	return w, nil
}

func (wtch *witches) WitchByID(ctx context.Context, uuid string) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.WitchByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return w, nil
}
func (wtch *witches) DeleteWitchByID(ctx context.Context, uuid string) error {
	err := wtch.witchesRepository.DeleteWitchByID(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (wtch *witches) UpdateWitchByID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.UpdateWitchByID(ctx, witch)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (wtch *witches) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	w, err := wtch.CreateWitch(ctx, witch)

	if err != nil {
		return nil, err
	}

	return w, nil
}
