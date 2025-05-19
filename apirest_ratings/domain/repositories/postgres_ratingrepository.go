package repository

import (
	"context"
	"database/sql"
	"fmt"

	"raitings.com/api/domain"
)

type PostgresRatingRepository struct {
	db *sql.DB
}

func NewPostgresRatingRepository(db *sql.DB) *PostgresRatingRepository {
	return &PostgresRatingRepository{
		db: db,
	}
}

func (r *PostgresRatingRepository) Save(ctx context.Context, rating *domain.Rating) error {
	query := "INSERT INTO ratings (movieid, userid, rating, timestamp) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query, rating.Movieid, rating.Userid, rating.Rating, rating.Timestamp)

	if err != nil {
		return fmt.Errorf("adapter: failed to save user %s to postgres: %w", rating.Movieid, err)
	}

	fmt.Printf("Adapter: Link %s saved to PostgreSQL.\n", rating.Movieid)
	return nil
}

func (r *PostgresRatingRepository) Close() error {
	return r.db.Close()
}
