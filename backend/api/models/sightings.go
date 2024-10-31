package models

import (
	"context"
	"database/sql"
)

type SQLSightingRepository struct {
	db *sql.DB
}

func NewSightingsRepo(db *sql.DB) *SQLSightingRepository {
	return &SQLSightingRepository{db: db}
}

func (r *SQLSightingRepository) AddSighting(ctx context.Context, sighting *Sighting) (uint64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `INSERT INTO sightings (latitude, longitude, city, area, date, time, situation, details, classification) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, query, sighting.Latitude, sighting.Longitude, sighting.City, sighting.Area, sighting.Date, sighting.Time, sighting.Situation, sighting.Details, sighting.Classification)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return uint64(id), nil

}
