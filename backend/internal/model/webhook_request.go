package model

type WebhookRequest struct {
	TenantID   string `json:"tenant_id"`
	ExternalID string `json:"external_id"`
	Message    string `json:"message"`
}
