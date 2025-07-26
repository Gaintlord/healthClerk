package models

import "time"

type Admin struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastLogin *time.Time `json:"last_login"`
}

type AdninLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLoginResponse struct {
	Token   string `json:"token"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
