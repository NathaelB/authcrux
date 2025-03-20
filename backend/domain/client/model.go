package client

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string    `json:"id"`
	Enabled   bool      `json:"enabled"`
	ClientID  string    `json:"client_id"`
	RealmID   string    `json:"realm_id"`
	Protocol  string    `json:"protocol"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewClient(realmID string, clientID string, name string) (*Client, error) {
	return &Client{
		ID:        uuid.New().String(),
		Enabled:   true,
		ClientID:  clientID,
		RealmID:   realmID,
		Protocol:  "openid-connect",
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
