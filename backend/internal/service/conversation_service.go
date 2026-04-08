package service

import "github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"

type ConversationService struct {
	repo repository.ConversationRepository
}

func NewConversationService(repo repository.ConversationRepository) *ConversationService {
	return &ConversationService{repo: repo}
}