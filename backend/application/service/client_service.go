package service

import "github.com/nathaelb/authcrux/domain/client"

type ClientService struct {
	clientRepo client.ClientRepository
}

func NewClientService(clientRepo client.ClientRepository) client.ClientService {
	return &ClientService{
		clientRepo: clientRepo,
	}
}

func (s *ClientService) CreateClient(realmID string, clientID string, name string) (*client.Client, error) {
	return s.clientRepo.Create(realmID, clientID, name)
}
