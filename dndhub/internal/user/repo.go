package user

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, username, passwordHash string) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx, `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
		RETURNING id, username, password_hash, created_at
	`, username, passwordHash).Scan(
		&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt,
	)
	return u, err
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, username, password_hash, created_at
		FROM users WHERE username = $1
	`, username).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt)
	return u, err
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, username, password_hash, created_at
		FROM users WHERE id = $1
	`, id).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.CreatedAt)
	return u, err
}
