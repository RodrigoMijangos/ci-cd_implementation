package repository

import (
	"context"
	"database/sql"
	"fmt"

	"raitings.com/api/domain"
)

type PostgresTagRepository struct {
	db *sql.DB
}

func NewPosgresTagRepository(db *sql.DB) *PostgresTagRepository {
	return &PostgresTagRepository{
		db: db,
	}
}

func (r *PostgresTagRepository) Save(ctx context.Context, tag *domain.Tag) error {
	query := "INSERT INTO tags (userid, movieid, tag, timestamp) VALUES ($1, $2, $3, $5)"

	_, err := r.db.ExecContext(ctx, query, tag.Userid, tag.Movieid, tag.Timestamp, tag.Tag)
	if err != nil {
		return fmt.Errorf("adapter: failed to save user %s to postgres: %w", tag.Tag, err)
	}
	fmt.Printf("Adapter: Tag %s saved to PostgreSQL.\n", tag.Tag)
	return nil
}
