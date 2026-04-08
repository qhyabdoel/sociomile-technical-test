package repository

import (
	"context"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type ConversationRepository interface {
	Create(ctx context.Context, conv *model.Conversation) error
	GetByID(ctx context.Context, tenantID int, id int) (*model.Conversation, error)
}