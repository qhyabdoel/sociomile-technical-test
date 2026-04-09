package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *model.Ticket) error
	UpdateStatus(ctx context.Context, tenantID, ticketID int64, status string) error
}

type ticketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ctx context.Context, ticket *model.Ticket) error {
	query := `
		INSERT INTO tickets (
			tenant_id, conversation_id, title, description, 
			status, priority, assigned_agent_id, created_at
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, NOW()
		)
	`
	_, err := r.db.ExecContext(ctx, query,
		ticket.TenantID,
		ticket.ConversationID,
		ticket.Title,
		ticket.Description,
		ticket.Status,
		ticket.Priority,
		ticket.AssignedAgentID,
	)
	return err
}

func (r *ticketRepository) UpdateStatus(ctx context.Context, tenantID, ticketID int64, status string) error {
	query := `UPDATE tickets SET status = ? WHERE tenant_id = ? AND id = ?`

	result, err := r.db.ExecContext(ctx, query, status, tenantID, ticketID)
	if err != nil {
		return err
	}

	// check if ticket is found
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("ticket not found")
	}

	return nil
}
