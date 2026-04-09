package repository

import (
	"context"
	"database/sql"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, tenant_id, name, email, password, role, created_at 
		FROM users 
		WHERE email = ?`

	user := &model.User{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found
		}
		return nil, err // some other error occurred
	}

	return user, nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	query := `
		SELECT id, tenant_id, role
		FROM users 
		WHERE id = ?`

	user := &model.User{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.TenantID,
		&user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no user found
		}
		return nil, err // some other error occurred
	}

	return user, nil
}
