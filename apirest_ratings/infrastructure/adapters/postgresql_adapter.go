package adapters

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(ctx context.Context, any struct{}) error {
	query := "INSERT INTO users (id, name) VALUES($1, $2);"

	_, err := r.db.ExecContext(ctx, query, any)

	if err != nil {
		return fmt.Errorf("adapter: failed to save user %s to postgres: %w", any, err)
	}
	fmt.Printf("Adapter: User %s saved to PostgreSQL.\n", any)

	return nil
}
