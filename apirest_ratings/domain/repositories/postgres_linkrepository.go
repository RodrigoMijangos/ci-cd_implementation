package repository

import (
	"context"
	"database/sql"
	"fmt"

	"raitings.com/api/domain"
)

type PostgresLinkRepository struct {
	db *sql.DB
}

func NewPostgresLinkRepository(db *sql.DB) *PostgresLinkRepository {
	return &PostgresLinkRepository{
		db: db,
	}
}

func (r *PostgresLinkRepository) Save(ctx context.Context, link *domain.Link) error {
	query := "INSERT INTO links (movieid, imbdid, tmbdid) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, link.Movieid, link.Imdbid, link.Tmdbid)

	if err != nil {
		return fmt.Errorf("adapter: failed to save user %d to postgres: %w", link.Movieid, err)
	}

	fmt.Printf("Adapter: Link %d saved to PostgreSQL.\n", link.Movieid)
	return nil
}
