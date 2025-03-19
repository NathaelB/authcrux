package service

import (
	"github.com/nathaelb/authcrux/domain/model"
	"github.com/nathaelb/authcrux/domain/port"
)

type RealmServiceImpl struct {
	realmRepo port.RealmRepository
}

func NewRealmService(realmRepo port.RealmRepository) port.RealmService {
	return &RealmServiceImpl{
		realmRepo: realmRepo,
	}
}

func (s *RealmServiceImpl) CreateRealm(name string) (*model.Realm, error) {
	realm, err := model.NewRealm(name)
	if err != nil {
		return nil, err
	}
	return s.realmRepo.Create(realm)
}

func (s *RealmServiceImpl) GetRealmByID(id string) (*model.Realm, error) {
	return s.realmRepo.GetByID(id)
}
