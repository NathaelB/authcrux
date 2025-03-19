package client

type ClientRepository interface {
	Create(client *Client) error
	GetByID(id string) (*Client, error)
	GetByClientID(clientID string) (*Client, error)
	Update(client *Client) error
	Delete(id string) error
}
