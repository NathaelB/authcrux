package realm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathaelb/authcrux/domain/port"
)

type RealmHandler struct {
	realmService port.RealmService
	handlers     struct {
		createRealm *CreateRealmHandler
	}
}

func NewRealmHandler(realmService port.RealmService) *RealmHandler {
	h := &RealmHandler{
		realmService: realmService,
	}

	h.handlers.createRealm = NewCreateRealmHandler(realmService)

	return h
}

func (h *RealmHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/realms")
	api.Post("/", h.handlers.createRealm.Handle)
}
