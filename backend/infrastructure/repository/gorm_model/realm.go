package gorm_model

import (
	"time"

	"github.com/nathaelb/authcrux/domain/model"
)

type Realm struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Realm) TableName() string {
	return "realms"
}

func (r *Realm) ToDomain() *model.Realm {
	return &model.Realm{
		ID:        r.ID,
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func RealmFromDomain(realm *model.Realm) *Realm {
	return &Realm{
		ID:        realm.ID,
		Name:      realm.Name,
		CreatedAt: realm.CreatedAt,
		UpdatedAt: realm.UpdatedAt,
	}
}
