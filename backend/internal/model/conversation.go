package model

import "time"

type Conversation struct {
	ID					int			`json:"id"`
	TenantID 			int			`json:"tenant_id"`
	CustomerExternalID 	string 		`json:"customer_external_id"`
	Status				string 		`json:"status"`
	AssignedAgentID		*int 		`json:"assigned_agent_id"`
	CreatedAt			time.Time 	`json:"created_at"`
	UpdatedAt			time.Time 	`json:"updated_at"`
}