package model

type User struct {
	ID           int64  `json:"id"`
	TenantID     int64  `json:"tenant_id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"` // "agent" or "customer"
}
