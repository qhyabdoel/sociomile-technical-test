package model

type User struct {
	ID       		int    `json:"id"`
	TenantID   		int    `json:"tenant_id"`
	Name       		string `json:"name"`
	Email      		string `json:"email"`
	PasswordHash 	string `json:"-"`
	Role	   		string `json:"role"` // "agent" or "customer"
}