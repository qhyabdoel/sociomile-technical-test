package model

type WebhookRequest struct {
	TenantID   int64  `json:"tenant_id"`
	ExternalID string `json:"external_id"`
	Message    string `json:"message"`
}
