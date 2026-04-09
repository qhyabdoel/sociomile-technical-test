package repository

import (
	"context"
	"database/sql"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type TenantRepository interface {
	GetByID(ctx context.Context, id int64) (*model.Tenant, error)
}

type tenantRepository struct {
	db *sql.DB
}

func NewTenantRepository(db *sql.DB) TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) GetByID(ctx context.Context, id int64) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := r.db.QueryRowContext(ctx, "SELECT * FROM tenants WHERE id = ?", id).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &tenant, nil
}
