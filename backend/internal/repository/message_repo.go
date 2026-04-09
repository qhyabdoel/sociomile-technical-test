package repository

import (
	"context"
	"database/sql"

	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
)

type MessageRepository interface {
	Create(ctx context.Context, msg *model.Message) error
	GetByConversationID(ctx context.Context, conversationID int64) ([]model.Message, error)
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(ctx context.Context, msg *model.Message) error {
	query := `
		INSERT INTO messages (conversation_id, sender_type, message) 
		VALUES (?, ?, ?)
	`
	_, err := r.db.ExecContext(ctx, query, msg.ConversationID, msg.SenderType, msg.Message)
	return err
}

func (r *messageRepository) GetByConversationID(ctx context.Context, conversationID int64) ([]model.Message, error) {
	query := `
		SELECT id, conversation_id, sender_type, message, created_at 
		FROM messages 
		WHERE conversation_id = ? 
		ORDER BY created_at ASC
	`
	rows, err := r.db.QueryContext(ctx, query, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(
			&msg.ID,
			&msg.ConversationID,
			&msg.SenderType,
			&msg.Message,
			&msg.CreatedAt,
		); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
