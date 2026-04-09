package service

import (
	"context"
	"errors"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"
)

type ConversationService struct {
	repo        repository.ConversationRepository
	messageRepo repository.MessageRepository
}

func NewConversationService(repo repository.ConversationRepository, msgRepo repository.MessageRepository) *ConversationService {
	return &ConversationService{repo: repo, messageRepo: msgRepo}
}

func (s *ConversationService) ProcessIncomingMessage(ctx context.Context, tenantID int64, externalID, message string) error {
	// get existing conversation
	conv, err := s.repo.FindByExternalID(ctx, tenantID, externalID)
	if err != nil {
		return err
	}

	// if not exist, create new conversation
	if conv == nil {
		conv = &model.Conversation{
			TenantID:           tenantID,
			CustomerExternalID: externalID,
			Status:             "open",
		}
		if err := s.repo.Create(ctx, conv); err != nil {
			return err
		}
	}

	// create message
	msg := &model.Message{
		ConversationID: conv.ID,
		SenderType:     "customer",
		Message:        message,
	}
	if err := s.repo.CreateMessage(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (s *ConversationService) GetConversationsByTenant(ctx context.Context, tenantID int64) ([]model.Conversation, error) {
	return s.repo.GetByTenant(ctx, tenantID)
}

// add agent reply to conversation
func (s *ConversationService) AddAgentReply(ctx context.Context, tenantID, convID int64, message string) error {
	conv, err := s.repo.GetByID(ctx, tenantID, convID)

	if err != nil || conv == nil {
		return errors.New("conversation not found")
	}

	if conv.Status == "closed" {
		return errors.New("conversation is closed")
	}

	msg := &model.Message{
		ConversationID: conv.ID,
		SenderType:     "agent",
		Message:        message,
	}

	if err := s.repo.CreateMessage(ctx, msg); err != nil {
		return err
	}

	return nil
}

func (s *ConversationService) GetConversationByID(ctx context.Context, tenantID, convID int64) (*model.ConversationDetail, error) {
	// validation
	conv, err := s.repo.GetByID(ctx, tenantID, convID)
	if err != nil || conv == nil {
		return nil, errors.New("conversation not found")
	}

	// get messages
	messages, err := s.messageRepo.GetByConversationID(ctx, convID)
	if err != nil {
		return nil, err
	}

	return &model.ConversationDetail{
		Conversation: *conv,
		Messages:     messages,
	}, nil
}
