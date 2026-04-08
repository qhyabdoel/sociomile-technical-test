package model

import "time"

type Ticket struct {
	ID             	int       `json:"id"`
	TenantID       	int       `json:"tenant_id"`
	ConversationID 	int       `json:"conversation_id"`
	Title		  	string    `json:"title"`
	Description     string    `json:"description"`
	Status         	string    `json:"status"` // "open", "pending", "closed"
	Priority       	string    `json:"priority"` // "low", "medium", "high"
	AssignedAgentID *int      `json:"assigned_agent_id,omitempty"` // nullable
	CreatedAt     	time.Time `json:"created_at"`
}