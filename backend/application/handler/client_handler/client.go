package client_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathaelb/authcrux/domain/client"
)

type ClientHandler struct {
	clientService client.ClientService
	handlers      struct {
		createClient *CreateClientHandler
	}
}

func NewClientHandler(clientService client.ClientService) *ClientHandler {
	h := &ClientHandler{
		clientService: clientService,
	}

	h.handlers.createClient = NewCreateClientHandler(clientService)

	return h
}

func (h *ClientHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/realms/:realm_id/clients")
	api.Post("/", h.handlers.createClient.Handle)
}
