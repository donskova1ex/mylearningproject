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

type Witches struct {
	db *sqlx.DB
}

func NewWitches(db *sqlx.DB) *Witches {
	return &Witches{db: db}
}

func (w *Witches) CreateWitch(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	var id uint32
	query := "INSERT INTO witches (uuid, name) values ($1, $2) RETURNING id" //TODO: обратботка на уникальный UUID во всех таблицах
	newUUID := uuid.NewString()
	row := w.db.QueryRow(query, newUUID, witch.Name)
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

func (w *Witches) WitchesAll(ctx context.Context) ([]*domain.Witch, error) {
	witches := []*domain.Witch{}
	rows, err := w.db.Queryx("SELECT uuid, id, name FROM witches")
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("empty table: %w", err)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&witches)
		if err != nil {
			return nil, fmt.Errorf("unable to perform witches select: %w", err)
		}

	}

	return witches, nil
}

func (w *Witches) WitchByUUID(ctx context.Context, uuid string) (*domain.Witch, error) {
	witch := &domain.Witch{}
	query := "SELECT id, name, uuid FROM witches WHERE uuid = $1"
	row := w.db.QueryRow(query, uuid) // TODO: поправить в остальных модулях
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, fmt.Errorf("witch with UUID: %s not found: %w", uuid, row.Err())
	}
	if row.Err() != nil {
		return nil, fmt.Errorf("can not read witch from db: %w", row.Err())
	}

	err := row.Scan(&witch)
	if err != nil {
		return nil, fmt.Errorf("can not create struct witch from db: %w", err)
	}
	return witch, nil

}

func (w *Witches) DeleteWitchByID(ctx context.Context, uuid string) error {
	_, err := w.db.Exec("DELETE FROM witches WHERE uuid = $1", uuid)
	if err != nil {
		return fmt.Errorf("can not delete witch with this ID: %w", err)
	}
	return nil
}

func (w *Witches) UpdateWitchByUUID(ctx context.Context, witch *domain.Witch) (*domain.Witch, error) {
	query := "UPDATE witches SET name = $1 WHERE uuid = $2"
	_, err := w.db.Exec(query, witch.Name, witch.UUID)
	if err != nil {
		return nil, fmt.Errorf("there is no object with this ID: %w", err)
	}
	return witch, nil
}
