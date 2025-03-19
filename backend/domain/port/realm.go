package port

import "github.com/nathaelb/authcrux/domain/model"

type RealmService interface {
	CreateRealm(name string) (*model.Realm, error)
	GetRealmByID(id string) (*model.Realm, error)
}

type RealmRepository interface {
	Create(realm *model.Realm) (*model.Realm, error)
	GetByID(id string) (*model.Realm, error)
}
