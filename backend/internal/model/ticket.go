package model

import "time"

type Ticket struct {
	ID              int64     `json:"id"`
	TenantID        int64     `json:"tenant_id"`
	ConversationID  int64     `json:"conversation_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`                      // "open", "pending", "closed"
	Priority        string    `json:"priority"`                    // "low", "medium", "high"
	AssignedAgentID *int64    `json:"assigned_agent_id,omitempty"` // nullable
	CreatedAt       time.Time `json:"created_at"`
}
