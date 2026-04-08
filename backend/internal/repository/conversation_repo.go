package repository

import (
	"context"
	"database/sql"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type ConversationRepository interface {
	FindByExternalID(ctx context.Context, tenantID int, externalID string) (*model.Conversation, error)
	Create(ctx context.Context, conv *model.Conversation) error
	GetByID(ctx context.Context, tenantID int, id int) (*model.Conversation, error)
}

type conversationRepo struct {
	db *sql.DB
}

// initiate repository with database connection
func NewConversationRepository(db *sql.DB) ConversationRepository {
	return &conversationRepo{db: db}
}

// find existing conversation by external ID
func (r *conversationRepo) FindByExternalID(ctx context.Context, tenantID int, externalID string) (*model.Conversation, error) {
	query := `
		SELECT id, tenant_id, customer_external_id, status, assigned_agent_id, created_at 
		FROM conversations 
		WHERE tenant_id = ? AND customer_external_id = ? 
		LIMIT 1`
	
	// this will hold the result of the query
	conv := &model.Conversation{}

	// execute the query and scan the result into the conversation struct
	err := r.db.QueryRowContext(ctx, query, tenantID, externalID).Scan(
		&conv.ID,
		&conv.TenantID,
		&conv.CustomerExternalID,
		&conv.Status,
		&conv.AssignedAgentID,
		&conv.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no conversation found
		}
		return nil, err // some other error occurred
	}

	return conv, nil
}

func (r *conversationRepo) Create(ctx context.Context, conv *model.Conversation) error {
	query := `
		INSERT INTO conversations (tenant_id, customer_external_id) 
		VALUES (?, ?, ?)`
	
	// execute the insert query
	_, err := r.db.ExecContext(ctx, query, 
		conv.TenantID,
		conv.CustomerExternalID, 
		conv.Status,
	)
	
	return err
}

func (r *conversationRepo) GetByID(ctx context.Context, tenantID int, id int) (*model.Conversation, error) {
	query := `
		SELECT id, tenant_id, customer_external_id, status, assigned_agent_id, created_at 
		FROM conversations 
		WHERE tenant_id = ? AND id = ?`
	
	conv := &model.Conversation{}

	err := r.db.QueryRowContext(ctx, query, tenantID, id).Scan(
		&conv.ID,
		&conv.TenantID,
		&conv.CustomerExternalID,
		&conv.Status,
		&conv.AssignedAgentID,
		&conv.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return conv, nil
}