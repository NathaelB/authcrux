package repository

import (
	"github.com/nathaelb/authcrux/domain/model"
	"github.com/nathaelb/authcrux/domain/port"
	"github.com/nathaelb/authcrux/infrastructure/repository/gorm_model"
	"gorm.io/gorm"
)

type PostgresRealmRepository struct {
	db *gorm.DB
}

func NewPostgresRealmRepository(db *gorm.DB) port.RealmRepository {
	return &PostgresRealmRepository{
		db: db,
	}
}

func (r *PostgresRealmRepository) Create(realm *model.Realm) (*model.Realm, error) {
	gormRealm := gorm_model.RealmFromDomain(realm)
	result := r.db.Create(gormRealm)
	if result.Error != nil {
		return nil, result.Error
	}
	return realm, nil
}

func (r *PostgresRealmRepository) GetByID(id string) (*model.Realm, error) {
	var gormRealm gorm_model.Realm
	result := r.db.First(&gormRealm, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return gormRealm.ToDomain(), nil
}
