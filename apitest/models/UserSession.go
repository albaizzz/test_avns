package models

import "time"

type UserSession struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Session     string    `json:"authorization"`
	CreatedDate string    `json:"created_date"`
	ExpiredDate time.Time `json:"expired_date"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	User
}
