package repository

import (
	"context"
	"database/sql"
	"fmt"

	"raitings.com/api/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (id, username, first_name, last_name, email, gender) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.ExecContext(
		ctx, query, user.Id, user.Username, user.First_name, user.Last_name, user.Email, user.Gender,
	)

	if err != nil {
		return fmt.Errorf("adapter: failed to save user %d to postgres: %w", user.Id, err)
	}

	fmt.Printf("Adapter: User %d saved to PostgreSQL.\n", user.Id)
	return nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	user := &domain.User{}
	err := row.Scan(&user.Id)

	if err == sql.ErrNoRows {
		return nil, nil // No encontrado
	}
	if err != nil {
		return nil, fmt.Errorf("adapter: failed to get user %d from postgres: %w", id, err)
	}

	return user, nil
}

func (r *PostgresUserRepository) Close() error {
	return r.db.Close()
}
