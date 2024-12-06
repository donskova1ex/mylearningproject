package processors

import (
	"context"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
)

type WitchesRepository interface {
	WitchesAll(ctx context.Context) ([]*domain.Witch, error)
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
