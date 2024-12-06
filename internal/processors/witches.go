package processors

import (
	"context"
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
	WitchesRepository WitchesRepository
	log               WitchesLogger
}
