package processors

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type WitchesRepository interface {
	CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error)
	WitchesAll(ctx context.Context) ([]*domain.Witch, error)
	WitchByUUID(ctx context.Context, uuid string) (*domain.Witch, error)
	DeleteWitchByUUID(ctx context.Context, uuid string) error
	UpdateWitchByUUID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error)
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
		wtch.log.Error("it is impossible to get a witches list", slog.String("err", err.Error()))
		return nil, fmt.Errorf("witches list processor error: %w", err)
	}
	return w, nil
}

func (wtch *witches) WitchByID(ctx context.Context, uuid string) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.WitchByUUID(ctx, uuid)
	if err != nil {
		wtch.log.Error("unable to get witch by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return nil, err
	}
	return w, nil
}
func (wtch *witches) DeleteWitchByID(ctx context.Context, uuid string) error {
	err := wtch.witchesRepository.DeleteWitchByUUID(ctx, uuid)
	if err != nil {
		wtch.log.Error("unable to delete witch by uuid",
			slog.String("err", err.Error()),
			slog.String("uuid", uuid))
		return err
	}
	return nil
}

func (wtch *witches) UpdateWitchByID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.UpdateWitchByUUID(ctx, witch)
	if err != nil {
		wtch.log.Error("unable to update witch by uuid")
		return nil, err
	}
	return w, nil
}

func (wtch *witches) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	w, err := wtch.witchesRepository.CreateWitch(ctx, witch)

	if err != nil {
		wtch.log.Error("unable to create witch",
			slog.String("err", err.Error()))
		return nil, err
	}

	return w, nil
}
