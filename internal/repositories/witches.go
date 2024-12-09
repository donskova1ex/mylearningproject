package repositories

import (
	"context"
	"fmt"

	"github.com/donskova1ex/mylearningproject/internal/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WitchesPostgres struct {
	db *sqlx.DB
}

func NewWitchesPostgres(db *sqlx.DB) *WitchesPostgres {
	return &WitchesPostgres{db: db}
}

func (w *WitchesPostgres) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	var id uint32
	query := "INCERT INTO witches (uuid, name) values ($1, $2) RETURNING id"
	newUUID := uuid.NewString()
	row := w.db.QueryRow(query, newUUID, witch.Name)
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("impossible to create an entity: %w", err)
	}
	newWitch := &domain.Witch{
		Name: witch.Name,
		UUID: newUUID,
		ID:   id,
	}
	return newWitch, nil
}
