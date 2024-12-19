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

type WitchesPostgres struct {
	db *sqlx.DB
}

func NewWitchesPostgres(db *sqlx.DB) *WitchesPostgres {
	return &WitchesPostgres{db: db}
}

func (w *WitchesPostgres) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	var id uint32

	query := "INSERT INTO witches (uuid, name) values ($1, $2) RETURNING id"

	newUUID := uuid.NewString()
	row := w.db.QueryRowContext(ctx, query, newUUID, witch.Name)
	err := row.Err()
	if err != nil {
		return nil, fmt.Errorf("can not read witch from db: %w", err)
	}
	if err := row.Scan(&id); err != nil {
		return nil, fmt.Errorf("can not scan witch for id: %w", err)
	}
	newWitch := &domain.Witch{
		Name: witch.Name,
		UUID: newUUID,
		ID:   id,
	}
	return newWitch, nil
}

func (w *WitchesPostgres) WitchesAll(ctx context.Context) ([]*domain.Witch, error) {
	witches := []*domain.Witch{}
	//rows, err := w.db.Queryx("SELECT uuid, id, name FROM witches")
	err := w.db.SelectContext(ctx, &witches, "SELECT uuid, id, name FROM witches")
	if errors.Is(err, sql.ErrNoRows) {
		return witches, nil
	}

	if err != nil {
		return nil, fmt.Errorf("can not read rows: %w", err)
	}
	return witches, nil
}

func (w *WitchesPostgres) WitchByUUID(ctx context.Context, uuid string) (*domain.Witch, error) {
	witch := &domain.Witch{}
	query := "SELECT id, name, uuid FROM witches WHERE uuid = $1"
	err := w.db.GetContext(ctx, witch, query, uuid) // TODO: поправить в остальных модулях
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("witch with UUID: %s not found: %w", uuid, err)
	}
	return witch, nil
}

func (w *WitchesPostgres) DeleteWitchByUUID(ctx context.Context, uuid string) error {
	_, err := w.db.ExecContext(ctx, "DELETE FROM witches WHERE uuid = $1", uuid)
	if err != nil {
		return fmt.Errorf("there is no object with this ID: %w", err)
	}
	return nil
}

func (w *WitchesPostgres) UpdateWitchByUUID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	query := "UPDATE witches SET name = $1 WHERE uuid = $2"
	_, err := w.db.Exec(query, witch.Name, witch.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return witch, nil
}
