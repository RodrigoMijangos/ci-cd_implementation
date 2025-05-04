package repository

import (
	"context"
	"database/sql"
	"fmt"

	"raitings.com/api/domain"
)

type PostgresMovieRepository struct {
	db *sql.DB
}

func NewPostgresMovieRepository(db *sql.DB) *PostgresMovieRepository {
	return &PostgresMovieRepository{
		db: db,
	}
}

func (r *PostgresMovieRepository) Save(ctx context.Context, movie *domain.Movie) error {
	query := "INSERT INTO movies (movieid, title, genres) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, movie.Movieid, movie.Title, movie.Genres)

	if err != nil {
		return fmt.Errorf("adapter: failed to save user %s to postgres: %w", movie.Movieid, err)
	}
	fmt.Printf("Adapter: User %s saved to PostgreSQL.\n", movie.Movieid)
	return nil
}

func (r *PostgresMovieRepository) Close() error {
	return r.db.Close()
}
