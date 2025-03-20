package client

type ClientService interface {
	CreateClient(realmID string, clientID string, name string) (*Client, error)
}

type ClientRepository interface {
	Create(realmID string, clientID string, name string) (*Client, error)
	GetByID(id string) (*Client, error)
	GetByClientID(clientID string) (*Client, error)
	Delete(id string) error
}
