package model

import "time"

type Message struct {
	ID             int       `json:"id"`
	ConversationID int       `json:"conversation_id"`
	Sender         string    `json:"sender"` // "customer" or "agent"
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}