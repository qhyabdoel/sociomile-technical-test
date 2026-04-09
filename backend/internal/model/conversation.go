package model

import "time"

type Conversation struct {
	ID                 int64     `json:"id"`
	TenantID           int64     `json:"tenant_id"`
	CustomerExternalID string    `json:"customer_external_id"`
	Status             string    `json:"status"`
	AssignedAgentID    *int64    `json:"assigned_agent_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ConversationDetail struct {
	Conversation Conversation `json:"conversation"`
	Messages     []Message    `json:"messages"`
}
