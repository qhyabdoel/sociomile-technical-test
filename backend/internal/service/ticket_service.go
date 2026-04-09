package service

import (
	"context"
	"errors"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"
)

type TicketService struct {
	ticketRepo repository.TicketRepository
	convRepo   repository.ConversationRepository
}

func NewTicketService(ticketRepo repository.TicketRepository, convRepo repository.ConversationRepository) *TicketService {
	return &TicketService{ticketRepo: ticketRepo, convRepo: convRepo}
}

// escalate conversation to ticket
func (s *TicketService) EscalateToTicket(ctx context.Context, tenantID, agentID, convID int64, title, desc, priority string) (*model.Ticket, error) {
	// validation
	conv, err := s.convRepo.GetByID(ctx, tenantID, convID)
	if err != nil || conv == nil {
		return nil, errors.New("conversation not found")
	}

	// create ticket
	ticket := &model.Ticket{
		TenantID:        tenantID,
		ConversationID:  convID,
		Title:           title,
		Description:     desc,
		Status:          "open",
		Priority:        priority,
		AssignedAgentID: &agentID,
	}

	// save ticket
	return s.ticketRepo.Create(ctx, ticket)
}

// update ticket status
func (s *TicketService) ChangeTicketStatus(ctx context.Context, tenantID, ticketID int64, newStatus string) error {
	// validation
	validStatus := map[string]bool{
		"open":        true,
		"in_progress": true,
		"resolved":    true,
		"closed":      true,
	}

	if !validStatus[newStatus] {
		return errors.New("invalid status")
	}

	// update ticket
	return s.ticketRepo.UpdateStatus(ctx, tenantID, ticketID, newStatus)
}
