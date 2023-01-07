package user

import (
	"time"
)

type User struct {
	Id           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	Address      string    `json:"address,omitempty"`
	Email        string    `json:"email,omitempty"`
	GeoLabel     string    `json:"geo_label,omitempty"`
	ChatId       string    `json:"chat_id,omitempty"`
	Password     string    `json:"password,omitempty"`
	RegisteredAt time.Time `json:"registered_at,omitempty"`
	LastLogin    time.Time `json:"last_login,omitempty"`
}
