package gorm_model

import (
	"time"

	"github.com/nathaelb/authcrux/domain/client"
)

type Client struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	Enabled   bool   `gorm:"not null"`
	ClientID  string `gorm:"not null"`
	RealmID   string `gorm:"not null"`
	Protocol  string `gorm:"not null"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Client) TableName() string {
	return "clients"
}

func (c *Client) ToDomain() *client.Client {
	return &client.Client{
		ID:        c.ID,
		Enabled:   c.Enabled,
		ClientID:  c.ClientID,
		RealmID:   c.RealmID,
		Protocol:  c.Protocol,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ClientFromDomain(client *client.Client) *Client {
	return &Client{
		ID:       client.ID,
		Enabled:  client.Enabled,
		ClientID: client.ClientID,
		RealmID:  client.RealmID,
		Protocol: client.Protocol,
		Name:     client.Name,
	}
}
