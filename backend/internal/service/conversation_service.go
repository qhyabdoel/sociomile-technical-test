package service

import (
	"context"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"
)

type ConversationService struct {
	repo repository.ConversationRepository
}

func NewConversationService(repo repository.ConversationRepository) *ConversationService {
	return &ConversationService{repo: repo}
}

func (s *ConversationService) ProcessIncomingMessage(ctx context.Context, tenantID, externalID, message string) error {
	return nil
}

func (s *ConversationService) GetConversationsByTenant(ctx context.Context, tenantID string) ([]model.Conversation, error) {
	return s.repo.GetConversationsByTenant(ctx, tenantID)
}

func (s *ConversationService) AddAgentReply(ctx context.Context, tenantID, convID, message string) error {
	return nil
}
