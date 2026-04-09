package repository

import (
	"context"
	"database/sql"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type ConversationRepository interface {
	FindByExternalID(ctx context.Context, tenantID int64, externalID string) (*model.Conversation, error)
	Create(ctx context.Context, conv *model.Conversation) error
	GetByID(ctx context.Context, tenantID int64, id int64) (*model.Conversation, error)
	GetByTenant(ctx context.Context, tenantID int64) ([]model.Conversation, error)
	CreateMessage(ctx context.Context, msg *model.Message) error
}

type conversationRepo struct {
	db *sql.DB
}

// initiate repository with database connection
func NewConversationRepository(db *sql.DB) ConversationRepository {
	return &conversationRepo{db: db}
}

func (r *conversationRepo) GetConversationsByTenant(ctx context.Context, tenantID int64) ([]model.Conversation, error) {
	query := `
		SELECT id, tenant_id, customer_external_id, status, assigned_agent_id, created_at 
		FROM conversations 
		WHERE tenant_id = ?`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversations []model.Conversation
	for rows.Next() {
		conv := model.Conversation{}
		err := rows.Scan(
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
		conversations = append(conversations, conv)
	}

	return conversations, nil
}

// find existing conversation by external ID
func (r *conversationRepo) FindByExternalID(ctx context.Context, tenantID int64, externalID string) (*model.Conversation, error) {
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

func (r *conversationRepo) GetByID(ctx context.Context, tenantID int64, id int64) (*model.Conversation, error) {
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

func (r *conversationRepo) CreateMessage(ctx context.Context, msg *model.Message) error {
	query := `
		INSERT INTO messages (conversation_id, sender_type, message) 
		VALUES (?, ?, ?)`

	// execute the insert query
	_, err := r.db.ExecContext(ctx, query,
		msg.ConversationID,
		msg.SenderType,
		msg.Message,
	)

	return err
}

func (r *conversationRepo) GetByTenant(ctx context.Context, tenantID int64) ([]model.Conversation, error) {
	query := `
		SELECT id, tenant_id, customer_external_id, status, assigned_agent_id, created_at 
		FROM conversations 
		WHERE tenant_id = ?`

	rows, err := r.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversations []model.Conversation
	for rows.Next() {
		conv := model.Conversation{}
		err := rows.Scan(
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
		conversations = append(conversations, conv)
	}

	return conversations, nil
}
