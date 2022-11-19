package models

type User struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	EmailConfirmed bool   `json:"email_confirmed"`
	IsAdmin        bool   `json:"is_admin"`
	IsSuperAdmin   bool   `json:"-"`
	OldPassword    string `json:"old_password"`
}
