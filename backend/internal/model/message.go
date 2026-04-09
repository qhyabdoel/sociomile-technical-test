package model

import "time"

type Message struct {
	ID             int64     `json:"id"`
	ConversationID int64     `json:"conversation_id"`
	SenderType     string    `json:"sender_type"` // "customer" or "agent"
	Message        string    `json:"message"`
	CreatedAt      time.Time `json:"created_at"`
}
