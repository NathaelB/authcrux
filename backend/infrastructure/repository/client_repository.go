package repository

import (
	"github.com/nathaelb/authcrux/domain/client"
	"github.com/nathaelb/authcrux/infrastructure/repository/gorm_model"
	"gorm.io/gorm"
)

type PostgresClientRepository struct {
	db *gorm.DB
}

func NewPostgresClientRepository(db *gorm.DB) client.ClientRepository {
	return &PostgresClientRepository{
		db: db,
	}
}

func (r *PostgresClientRepository) Create(realmID string, clientID string, name string) (*client.Client, error) {
	client, err := client.NewClient(realmID, clientID, name)
	if err != nil {
		return nil, err
	}

	gormClient := gorm_model.ClientFromDomain(client)
	result := r.db.Create(gormClient)
	if result.Error != nil {
		return nil, result.Error
	}

	return client, nil
}

func (r *PostgresClientRepository) GetByID(id string) (*client.Client, error) {
	var gormClient gorm_model.Client
	result := r.db.First(&gormClient, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return gormClient.ToDomain(), nil
}

func (r *PostgresClientRepository) GetByClientID(clientID string) (*client.Client, error) {
	var gormClient gorm_model.Client
	result := r.db.First(&gormClient, "client_id = ?", clientID)
	if result.Error != nil {
		return nil, result.Error
	}

	return gormClient.ToDomain(), nil
}

func (r *PostgresClientRepository) Delete(id string) error {
	result := r.db.Delete(&gorm_model.Client{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
